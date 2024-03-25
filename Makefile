migrateup:
	migrate -database "postgres://$(shell echo $$DB_USERNAME):$(shell echo $$DB_PASSWORD)@$(shell echo $$DB_HOST):$(shell echo $$DB_PORT)/$(shell echo $$DB_NAME)?$(shell echo $$DB_PARAMS)" -path db/migrations up

migratedown:
	migrate -database "postgres://$(shell echo $$DB_USERNAME):$(shell echo $$DB_PASSWORD)@$(shell echo $$DB_HOST):$(shell echo $$DB_PORT)/$(shell echo $$DB_NAME)?$(shell echo $$DB_PARAMS)" -path db/migrations down

run:
	go run main.go