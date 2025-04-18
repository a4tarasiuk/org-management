package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"management/internal/app/models"
)

var ormModels = []any{
	&models.Organization{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(ormModels...)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}
