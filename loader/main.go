package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"management/internal/app/orm"
)

var models = []any{
	&orm.Organization{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(models...)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}

	io.WriteString(os.Stdout, stmts)
}
