package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

// runs

func (store *Store) GetRuns(fromID int64) []float64 {
	/*out := make([]float64, len(state.runs[fromId]))
	copy(out, state.runs[fromId])
	return out*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rows, err := store.pool.Query(ctx, "SELECT distance FROM runs WHERE user_id = $1 ORDER BY id ASC",
		fromID,
	)
	if err != nil {
		log.Println(err)
		return []float64{}
	}

	defer rows.Close()

	out := make([]float64, 0)

	for rows.Next() {
		var f float64

		if err := rows.Scan(&f); err != nil {
			log.Println(err)
			return []float64{}
		}

		out = append(out, f)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return []float64{}
	}

	return out
}

func (store *Store) SetWaitingDistance(fromID int64) {
	//state.waitingDistance[fromId] = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := store.pool.Exec(ctx, "INSERT INTO users(user_id, waiting_distance) VALUES ($1, true) "+
		"ON CONFLICT (user_id) DO UPDATE SET waiting_distance = EXCLUDED.waiting_distance",
		fromID,
	)

	if err != nil {
		log.Println(err)
	}

}

func (store *Store) ClearWaitingDistance(fromID int64) {
	//delete(state.waitingDistance, fromId)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := store.pool.Exec(ctx, "UPDATE users SET waiting_distance = $1 WHERE user_id = $2",
		false,
		fromID,
	)

	if err != nil {
		log.Println(err)
	}

}

func (store *Store) AddRun(fromID int64, dist float64) {
	//state.runs[fromId] = append(state.runs[fromId], dist)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := store.pool.Exec(ctx, "INSERT INTO runs(user_id, distance) VALUES ($1, $2)",
		fromID,
		dist,
	)

	if err != nil {
		log.Println(err)
	}
}

// users

func (store *Store) SetWaitingName(fromID int64) {
	//state.waitingName[fromId] = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := store.pool.Exec(ctx, "INSERT INTO users(user_id, waiting_name) VALUES ($1, true) "+
		"ON CONFLICT (user_id) DO UPDATE SET waiting_name = EXCLUDED.waiting_name",
		fromID,
	)

	if err != nil {
		log.Println(err)
	}
}

func (store *Store) AddName(fromID int64, name string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := store.pool.Exec(ctx, "INSERT INTO users(user_id, name) VALUES ($1, $2) "+
		"ON CONFLICT (user_id) DO UPDATE SET name = EXCLUDED.name",
		fromID,
		name,
	)

	if err != nil {
		log.Println(err)
	}
}

func (store *Store) ClearWaitingName(fromID int64) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := store.pool.Exec(ctx, "UPDATE users SET waiting_name = $1 WHERE user_id = $2",
		false,
		fromID,
	)

	if err != nil {
		log.Println(err)
	}
}

func (store *Store) GetName(fromID int64) string {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var name string

	err := store.pool.QueryRow(ctx, "SELECT name FROM users WHERE user_id = $1",
		fromID,
	).Scan(&name)

	if err != nil {
		log.Println(err)
		return ""
	}

	return name
}

// checks
func (store *Store) IsRegistered(fromID int64) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var exists bool
	err := store.pool.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE user_id = $1 AND name <> '')",
		fromID,
	).Scan(&exists)

	if err != nil {
		log.Println(err)
		return false
	}

	return exists
}

func (store *Store) IsWaitingDistance(fromID int64) bool {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var exists bool
	err := store.pool.QueryRow(ctx, "SELECT waiting_distance FROM users WHERE user_id = $1",
		fromID,
	).Scan(&exists)

	if err != nil {
		return false
	}

	return exists
}

func (store *Store) IsWaitingName(fromID int64) bool {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var exists bool

	err := store.pool.QueryRow(ctx, "SELECT waiting_name FROM users WHERE user_id = $1",
		fromID,
	).Scan(&exists)

	if err != nil {
		return false
	}

	return exists
}
