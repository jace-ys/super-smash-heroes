.PHONY: integration setup teardown proto migrations format

PROTO = v1/battle v1/superhero

integration:
	@make teardown
	@make setup
	@echo "==> Initializing.." && ../bin/wait-for localhost:5432
	@echo "==> Running integration tests.."
	go test -v ./integration

images:
	@echo "==> Building images.."
	docker-compose build

setup:
	@echo "==> Starting containers.."
	docker-compose -f ../docker-compose.yaml -f docker-compose.test.yaml up -d

teardown:
	@echo "==> Killing containers.."
	docker-compose down -v

proto:
	@echo "==> Generating proto.."
	../bin/protogen -t "${PROTO}" -p ../api/proto

migrations:
	@echo "==> Copying migrations.."
	rsync -v --ignore-existing ../databases/postgres/migrations/ data/postgres/migrations/

format:
	@echo "==> Formatting code.."
	gofmt -w .
