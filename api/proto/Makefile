SERVICES = battle superhero
GODIR = generated/go

.PHONY: all goproto

all: goproto

goproto:
	@echo "==> Compiling proto files"
	@for s in ${SERVICES}; do \
		mkdir -p ${GODIR}/$$s; \
		protoc -I=services/$$s --go_out=plugins=grpc:${GODIR}/$$s services/$$s/*.proto; \
		echo "Done compiling .proto files for service: $$s"; \
	done
