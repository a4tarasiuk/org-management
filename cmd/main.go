package main

import (
	"log"

	"management/internal/infra/conf"
)

func main() {
	config, err := conf.Load(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
}
