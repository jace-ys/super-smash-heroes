.PHONY: integration images setup teardown tools proto migrations format

PROTO = v1/battle v1/superhero

integration:
	@make teardown
	@echo "==> Initializing.."
	@make setup
	@echo "==> Running integration tests.."
	go test -v ./integration

images:
	@echo "==> Building images.."
	docker-compose build

setup:
	@echo "==> Starting containers.."
	docker-compose -f ../docker-compose.yaml -f docker-compose.test.yaml up -d
	../bin/wait-for localhost:5001 && sleep 10

teardown:
	@echo "==> Killing containers.."
	docker-compose down -v

tools:
	@echo "==> Installing tools.."
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

proto:
	@echo "==> Generating proto.."
	../bin/protogen -t "${PROTO}" -p ../api/proto

migrations:
	@echo "==> Copying migrations.."
	rsync -v --ignore-existing ../databases/postgres/migrations/ data/postgres/migrations/

format:
	@echo "==> Formatting code.."
	go fmt ./...
