package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"management/internal/infra/conf"
)

func InitDB(c conf.Config) (*gorm.DB, error) {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.PgUser, c.PgPass, c.PgHost, c.PgPort, c.PgDB)

	return gorm.Open(postgres.Open(connStr))
}
