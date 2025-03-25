include .env

MIGRATIONS_PATH = ./migrations

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make create name=<filename>"; \
		exit 1; \
	fi
	@echo "Creating migration files for $(name)"
	@migrate create -seq -ext .sql -dir ${MIGRATIONS_PATH} $(name)

.PHONY: migrate-up
migrate-up:
	@echo "Running up migrations"
	@migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} up

.PHONY: migrate-down
migrate-down:
	@echo "Running down migrations"
	@migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} down

.PHONY: migrate-drop
migrate-drop:
	@echo "Dropping all tables"
	@migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} drop