.PHONY: all build run proto dependencies test format

include ../../.env

SOURCE = cmd/service/main.go
PROTO = v1/superhero
DEPENDENCIES = db.postgres db.postgres.init

all: format run

build:
	@echo "==> Compiling code.."
	go build -o service ${SOURCE}

run:
	@echo "==> Executing code.."
	go run ${SOURCE} \
	--port 8081 \
	--gateway-port 8080 \
	--postgres-host 127.0.0.1:5432 \
	--postgres-user nintendo \
	--postgres-password nintendo \
	--postgres-db nintendo \
	--superhero-api-token ${SUPERHERO_API_TOKEN}

proto:
	@echo "==> Generating proto.."
	../../bin/protogen -p ../../api/proto -t "${PROTO}" -g

dependencies:
	@echo "==> Starting auxiliary containers.."
	docker-compose up -d ${DEPENDENCIES}

test:
	@echo "==> Running tests.."
	go test -v ./...

format:
	@echo "==> Formatting code.."
	gofmt -w .
