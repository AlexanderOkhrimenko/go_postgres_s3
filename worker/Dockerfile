FROM golang:latest AS builder
LABEL maintainer="Aleksandr Okhrimenko"

RUN mkdir /go/src/go_postgres_s3/
WORKDIR /go/src/go_postgres_s3/worker/

RUN go get "github.com/lib/pq" && \
    go get "github.com/minio/minio-go"

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o worker .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/go_postgres_s3/worker .

EXPOSE 8080
CMD ["./worker"]