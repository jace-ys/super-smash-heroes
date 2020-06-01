.PHONY: postgres

postgres:
	@echo "==> Generating new migration files for Postgres.."
	migrate create -ext sql -dir postgres/migrations ${NAME}
