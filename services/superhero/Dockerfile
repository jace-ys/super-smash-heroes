FROM alpine:3.7
RUN apk add --no-cache ca-certificates bash
WORKDIR /usr/src/app
COPY main_docker .
EXPOSE 80
CMD ["./main_docker"]
