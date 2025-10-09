DB_NAME=chatappdb
DB_HOST=localhost:5432
DB_USER=postgres
DB_PASSWORD=root
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)?sslmode=disable

MIGRATION_DIR=./migrations

migrate-create:
	migrate create  -dir $(MIGRATION_DIR) -ext sql -seq $(MG_NAME)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database $(DB_URL) up 

migrate-down:
	migrate -path $(MIGRATION_DIR) -database $(DB_URL) down

migrate-rollback:
	migrate -database $(DB_URL) -path $(MIGRATION_DIR) force $(VER)