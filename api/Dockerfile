FROM golang:latest AS builder
LABEL maintainer="Aleksandr Okhrimenko"

RUN mkdir /go/src/go_postgres_s3/
WORKDIR /go/src/go_postgres_s3/api/

RUN go get "github.com/lib/pq" && \
    go get "github.com/gin-gonic/gin"

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/go_postgres_s3/api .

EXPOSE 8080
CMD ["./api"]
