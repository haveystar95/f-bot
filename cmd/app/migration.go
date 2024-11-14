package main

import (
	"database/sql"
	"f-bot/pkg/config"
	"f-bot/pkg/db"
	"f-bot/pkg/logger"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

func runMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not run up migrations: %v", err)
	}

	fmt.Println("Migrations ran successfully")
}

func main() {
	err := godotenv.Load()

	cfg := config.LoadConfig()
	logger.InitLogger()
	db.InitDB(cfg.Database.Host, strconv.Itoa(cfg.Database.Port), cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	sq, err := db.DB.DB()

	if err != nil {
		log.Fatalf("Could not get sql.DB instance: %v", err)
	}

	runMigration(sq)

}
