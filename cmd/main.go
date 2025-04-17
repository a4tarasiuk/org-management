package main

import (
	"log"

	"management/internal/app"
	"management/internal/infra"
	"management/internal/infra/conf"
)

func main() {
	config, err := conf.Load(".")
	if err != nil {
		log.Fatalf("Cannot load config: %s", err)
	}

	_infra, err := infra.Init(config)

	if err != nil {
		log.Fatalf("Cannot init the application: %s", err)
	}

	app := app.Application{}
	app.Init(_infra)

	if err = _infra.Run(); err != nil {
		log.Fatalf("Unable to run the application: %s", err)
	}
}
