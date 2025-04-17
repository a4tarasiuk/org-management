.PHONY: create-migrations apply-migrations

include .env
export

create-migrations:
		export PG_CONN_STR="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

		@read -p "Enter migration name: " migration_name; \
		atlas migrate diff $$migration_name --env gorm

apply-migrations:
		export PG_CONN_STR="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"

		atlas migrate apply --env gorm
