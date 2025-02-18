# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN go build -o websocket-manager cmd/main.go

# Expose port
EXPOSE 8080

# Run the application
CMD ["./websocket-manager"]
