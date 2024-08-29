FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o hotel-service cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates bash curl

WORKDIR /root/

COPY --from=builder /app/hotel-service .


EXPOSE 8081

CMD ["./hotel-service", "--config", "/cmd/configs/config.yaml"]
