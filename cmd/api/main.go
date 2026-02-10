package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DiMashina05/tg-running-tracker/internal/httpapi"
	"github.com/DiMashina05/tg-running-tracker/internal/storage"
	"github.com/DiMashina05/tg-running-tracker/internal/storage/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main(){
	dsn := os.Getenv("DATABASE_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Panic(err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Panic(err)
	}

	var store storage.Store = postgres.New(pool)

	handler := httpapi.NewServer(store)

	port := os.Getenv("SERVER_PORT")

	if err := http.ListenAndServe(port, handler); err != nil{
		log.Fatal(err)
	}	
}
