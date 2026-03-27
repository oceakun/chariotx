FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o map-service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/map-service .
EXPOSE 8081
CMD ["./map-service"]
