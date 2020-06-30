FROM golang:latest AS builder

RUN mkdir /go/src/go_postgres_s3/
WORKDIR /go/src/go_postgres_s3/

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/go_postgres_s3/ .

EXPOSE 8080
CMD ["./api"]