FROM golang:1.14 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix cgo ./cmd/service/...

FROM alpine:3.11
WORKDIR /src
COPY --from=builder /go/bin/ /bin/
CMD ["service"]
