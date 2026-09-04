package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/zahidhasanzh/olx-api/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up | down>")
	}

	cnf := config.MustLoad()

	m, err := migrate.New(
		"file://migrations", cnf.DatabaseUrl)
	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
		}
	case "clean":
		if err := m.Drop(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}

}
