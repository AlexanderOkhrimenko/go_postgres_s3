FROM golang:latest AS builder
LABEL maintainer="Aleksander Okhrimenko"

RUN mkdir /go/src/worker
WORKDIR /go/src/worker/

RUN go get "github.com/lib/pq" && \
    go get "github.com/minio/minio-go"

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o worker .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/worker .

EXPOSE 8080
CMD ["./worker"]