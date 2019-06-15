TARGET = main
SOURCE = main.go
IMAGE = jaceys/super-smash-heroes/service.superhero
VERSION = latest

.PHONY: all run build execute image test format clean

all: format build execute

run:
	@echo "==> Running source code.."
	go run ${SOURCE}

build:
	@echo "==> Building from source.."
	go build -o ${TARGET} ${SOURCE}

execute:
	@echo "==> Executing binary.."
	@./${TARGET}

image:
	@echo "==> Building Docker image"
	docker build -t ${IMAGE}:${VERSION} .

test:
	@echo "==> Running tests.."
	go test ./... -v

format:
	@echo "==> Formatting code.."
	gofmt -w .

clean:
	@echo "==> Cleaning up.."
	go mod tidy
	go clean
