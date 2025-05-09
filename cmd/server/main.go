package main

import (
	"aarav345/RiderAppGoLang.git/config"
	"aarav345/RiderAppGoLang.git/db"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := db.NewPostgresPool(config.Envs.Postgres)

	if err != nil {
		log.Fatal("Failed to connect using pgxpool:", err)
	}
	defer pool.Close()
	initStorage(pool)
}

func initStorage(pool *pgxpool.Pool) {
	err := pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}
	log.Println("Connected to the database successfully")
}
