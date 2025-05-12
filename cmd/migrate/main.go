package main

import (
	"aarav345/RiderAppGoLang.git/config"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	postgresURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Envs.Postgres.User,
		config.Envs.Postgres.Password,
		config.Envs.Postgres.Host,
		config.Envs.Postgres.Port,
		config.Envs.Postgres.DBName,
		config.Envs.Postgres.SSLMode,
	)

	m, err := migrate.New(
		"file://cmd/migrate/migrations",
		postgresURL,
	)

	if err != nil {
		log.Fatal("Failed to intialize migration:", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("Missing migration command (up or down)")
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration up failed: %v", err)
		}
		fmt.Println("Migration up successful")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration down failed: %v", err)
		}
		fmt.Println("Migration down successful")

	default:
		log.Fatalf("Unknown command %q", os.Args[1])
	}

}
