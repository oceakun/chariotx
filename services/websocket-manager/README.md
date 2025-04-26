# 📍 Location Service

## ⚙️ Setup Instructions

### 1. Run Docker Compose

```bash
cd ./services/websocket-manager
docker-compose up -d
```

### 2. Run Server

```bash
go mod tidy
go run cmd/main.go
```

---

## 🧪 Testing Instructions

### 1. Run Client

```bash
go run tests/integration/user_ws_connection.go
```

### 2. Verify

#### ✅ Check Redis

```bash
docker exec -it cassandra cqlsh -e "SELECT * FROM location.locations;"
```
