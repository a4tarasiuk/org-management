data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

variable "postgres_connection_string" {
  type    = string
  default = getenv("PG_CONN_STR")
}

env "gorm" {
  src = data.external_schema.gorm.url

  url = var.postgres_connection_string

  dev = "docker://postgres/15/dev?search_path=public"

  migration {
    dir = "file://internal/infra/postgres/migrations"

    format = golang-migrate
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
