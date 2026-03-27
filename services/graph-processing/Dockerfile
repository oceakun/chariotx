FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o graph-processing .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/graph-processing .
EXPOSE 8080
CMD ["./graph-processing"]
