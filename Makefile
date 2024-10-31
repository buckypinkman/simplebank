ifneq (,$(wildcard ./.env))
    include .env
    export
endif

MIGRATION_DIR=db/migration

postgres:
	docker run --name my-postgres -p 5432:${DB_PORT} -e POSTGRES_USER=${DB_USERNAME} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres

createdb:
	docker exec -it my-postgres createdb --username=${DB_USERNAME} --owner=${DB_USERNAME} ${DB_NAME}

dropdb:
	docker exec -it my-postgres dropdb --username=${DB_USERNAME} ${DB_NAME}

migrateup:
	migrate -path $(MIGRATION_DIR) -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable" -verbose up 

migratedown:
	migrate -path $(MIGRATION_DIR) -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migration migrateup migratedown sqlc test

migration:
ifndef name
	$(error name is not set. Usage: make migration name=<migration_name>)
endif
	@echo "Creating new migration: $(name)"
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)