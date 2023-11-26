FROM golang:1.21.4 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o dist/ ./cmd/service/...

FROM alpine:3.18
WORKDIR /src
COPY --from=builder /src/dist /bin/
CMD ["/bin/service"]
