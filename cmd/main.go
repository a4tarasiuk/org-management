package main

import (
	"log"

	"management/internal/infra"
	"management/internal/infra/conf"
)

func main() {
	config, err := conf.Load(".")
	if err != nil {
		log.Fatalf("Cannot load config: %s", err)
	}

	app, err := infra.Init(config)

	if err != nil {
		log.Fatalf("Cannot init the application: %s", err)
	}
}
