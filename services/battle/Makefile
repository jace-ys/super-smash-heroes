.PHONY: all run build dependencies tools proto test format

SOURCE = cmd/service/main.go
PROTO = v1/battle
DEPENDENCIES = db.postgres db.postgres.init

all: format run

run:
	@echo "==> Executing code.."
	@go run ${SOURCE} \
		--port 8081 \
		--gateway-port 8080 \
		--database-url postgres://nintendo:nintendo@127.0.0.1:5432/nintendo?sslmode=disable

build:
	@echo "==> Compiling code.."
	go build -o service ${SOURCE}

dependencies:
	@echo "==> Starting auxiliary containers.."
	docker compose up -d ${DEPENDENCIES}

tools:
	@echo "==> Installing tools.."
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

proto:
	@echo "==> Generating proto.."
	../../bin/protogen -p ../../api/proto -t "${PROTO}" -g

test:
	@echo "==> Running tests.."
	go test -v ./...

format:
	@echo "==> Formatting code.."
	go fmt ./...
