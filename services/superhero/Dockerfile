FROM golang:1.12 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main

FROM alpine:3.7
WORKDIR /usr/src/app
COPY --from=builder /build/main .
EXPOSE 80
CMD ["./main"]
