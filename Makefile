migrate-up:
		migrate -path=internal/infra/db/migrations \
				-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
				-verbose up

migrate-down:
		migrate -path=internal/infra/db/migrations \
				-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
				-verbose down
