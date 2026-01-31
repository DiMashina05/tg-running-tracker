//Вайб код тесты

package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func testDBURL() string {
	if v := os.Getenv("TEST_DATABASE_URL"); v != "" {
		return v
	}
	return "postgres://rt_user:rt_pass@127.0.0.1:5434/running_tracker_test"
}

func newTestPool(t *testing.T) *pgxpool.Pool {
	t.Helper()

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, testDBURL())
	if err != nil {
		t.Fatalf("pgxpool.New: %v", err)
	}

	t.Cleanup(func() { pool.Close() })
	return pool
}

func cleanDB(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()

	ctx := context.Background()

	_, err := pool.Exec(ctx, "TRUNCATE runs, users RESTART IDENTITY")
	if err != nil {
		t.Fatalf("TRUNCATE: %v", err)
	}
}

func TestSetWaitingName_CreatesUserAndSetsFlag(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(111)

	store.SetWaitingName(userID)

	ctx := context.Background()
	var waiting bool
	err := pool.QueryRow(ctx, "SELECT waiting_name FROM users WHERE user_id = $1", userID).Scan(&waiting)
	if err != nil {
		t.Fatalf("select waiting_name: %v", err)
	}
	if !waiting {
		t.Fatalf("expected waiting_name=true, got false")
	}
}

func TestAddName_MakesUserRegistered(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(222)

	store.AddName(userID, "Dima")

	if !store.IsRegistered(userID) {
		t.Fatalf("expected IsRegistered=true after AddName")
	}

	ctx := context.Background()
	var name string
	err := pool.QueryRow(ctx, "SELECT name FROM users WHERE user_id = $1", userID).Scan(&name)
	if err != nil {
		t.Fatalf("select name: %v", err)
	}
	if name != "Dima" {
		t.Fatalf("expected name=Dima, got %q", name)
	}
}

func TestAddRun_ThenGetRuns_ReturnsInOrder(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(333)

	store.AddRun(userID, 5)
	store.AddRun(userID, 10)

	runs := store.GetRuns(userID)
	if len(runs) != 2 {
		t.Fatalf("expected 2 runs, got %d (%v)", len(runs), runs)
	}
	if runs[0] != 5 || runs[1] != 10 {
		t.Fatalf("expected [5 10], got %v", runs)
	}
}

func TestClearWaitingName_SetsFlagFalse(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(444)

	store.SetWaitingName(userID)
	store.ClearWaitingName(userID)

	ctx := context.Background()
	var waiting bool
	err := pool.QueryRow(ctx, "SELECT waiting_name FROM users WHERE user_id = $1", userID).Scan(&waiting)
	if err != nil {
		t.Fatalf("select waiting_name: %v", err)
	}
	if waiting {
		t.Fatalf("expected waiting_name=false, got true")
	}
}

func TestIsWaitingName_ReflectsDBState(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(777)

	store.SetWaitingName(userID)
	if !store.IsWaitingName(userID) {
		t.Fatalf("expected IsWaitingName=true after SetWaitingName")
	}

	store.ClearWaitingName(userID)
	if store.IsWaitingName(userID) {
		t.Fatalf("expected IsWaitingName=false after ClearWaitingName")
	}
}


func TestSetWaitingDistance_CreatesUserAndSetsFlag(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(555)

	store.SetWaitingDistance(userID)

	ctx := context.Background()
	var waiting bool
	err := pool.QueryRow(ctx, "SELECT waiting_distance FROM users WHERE user_id = $1", userID).Scan(&waiting)
	if err != nil {
		t.Fatalf("select waiting_distance: %v", err)
	}
	if !waiting {
		t.Fatalf("expected waiting_distance=true, got false")
	}
}

func TestIsWaitingDistance_ReflectsDBState(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(888)

	store.SetWaitingDistance(userID)
	if !store.IsWaitingDistance(userID) {
		t.Fatalf("expected IsWaitingDistance=true after SetWaitingDistance")
	}

	store.ClearWaitingDistance(userID)
	if store.IsWaitingDistance(userID) {
		t.Fatalf("expected IsWaitingDistance=false after ClearWaitingDistance")
	}
}


func TestClearWaitingDistance_SetsFlagFalse(t *testing.T) {
	pool := newTestPool(t)
	cleanDB(t, pool)

	store := New(pool)
	userID := int64(666)

	store.SetWaitingDistance(userID)
	store.ClearWaitingDistance(userID)

	ctx := context.Background()
	var waiting bool
	err := pool.QueryRow(ctx, "SELECT waiting_distance FROM users WHERE user_id = $1", userID).Scan(&waiting)
	if err != nil {
		t.Fatalf("select waiting_distance: %v", err)
	}
	if waiting {
		t.Fatalf("expected waiting_distance=false, got true")
	}
}

