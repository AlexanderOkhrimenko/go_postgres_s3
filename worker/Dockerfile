FROM golang:latest AS builder

RUN mkdir /go/src/worker
WORKDIR /go/src/worker/

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o worker .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/worker .

EXPOSE 8080
CMD ["./worker"]