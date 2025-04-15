package main

import (
	"log"

	"management/internal/infra/conf"
	"management/internal/infra/postgres"
)

func main() {
	config, err := conf.Load(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	db, err := postgres.InitDB(config)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
}
