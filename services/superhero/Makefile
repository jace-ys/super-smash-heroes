TARGET = main
SOURCE = main.go
DOCKER_BINARY = main_docker
IMAGE = super-smash-heroes/service.superhero
VERSION = latest
CONFIG_FILE = ../../config/config.dev.yaml

.PHONY: all run build execute image test format update

all: format build execute

run:
	@echo "==> Running source code.."
	go run ${SOURCE}

build:
	@echo "==> Building from source.."
	go build -o ${TARGET} ${SOURCE}

execute:
	@echo "==> Executing binary.."
	CONFIG_FILE=${CONFIG_FILE} ./${TARGET}

image:
	@echo "==> Building binary for Linux.."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${DOCKER_BINARY}
	@echo "==> Building image.."
	docker build -t ${IMAGE}:${VERSION} .
	@echo "==> Removing binary.."
	rm ${DOCKER_BINARY}

test:
	@echo "==> Running tests.."
	go test ./... -v

format:
	@echo "==> Formatting code.."
	gofmt -w .

update:
	@echo "==> Updating dependencies.."
	go get -u
	go mod tidy
