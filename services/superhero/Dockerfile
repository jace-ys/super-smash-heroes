FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /usr/src/app
COPY main_docker .
EXPOSE 80
CMD ["./main_docker"]
