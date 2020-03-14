.PHONY: all build run proto dependencies test format

SOURCE = cmd/service/main.go
PROTO = v1/superhero
DEPENDENCIES = db.postgres

all: format run

build:
	@echo "==> Compiling code.."
	go build ${SOURCE}

run:
	@echo "==> Executing code.."
	go run ${SOURCE}

proto:
	@echo "==> Generating proto.."
	../../bin/protogen -p ../../api/proto -t "${PROTO}"

dependencies:
	@echo "==> Starting auxiliary containers.."
	docker-compose up -d ${DEPENDENCIES}

test:
	@echo "==> Running tests.."
	go test -v ./...

format:
	@echo "==> Formatting code.."
	gofmt -w .
