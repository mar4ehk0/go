FROM golang:1.22.6 as builder

WORKDIR /build/shop/
COPY ../../.dockerignore ./
COPY ../../go.mod ../../go.sum ./

RUN go mod download

COPY ../ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shop_api ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build/shop/shop_api ./
COPY --from=builder /build/shop/.env ./
EXPOSE 8080
CMD ["./shop_api", "-p=8080", "-i=0.0.0.0"]