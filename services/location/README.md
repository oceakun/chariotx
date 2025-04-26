# 📍 Location Service

## ⚙️ Setup Instructions

### 1. Run Docker Compose

```bash
cd ./services/location
docker-compose up -d
```

### 2. Initialize Cassandra

```bash
chmod +x cassandra-init.sh
./cassandra-init.sh
```

### 3. Initialize Kafka

```bash
chmod +x kafka-init.sh
./kafka-init.sh
```

### 4. Run Server

```bash
go mod tidy
go run main.go
```

---

## 🧪 Testing Instructions

### 1. Run Client

```bash
go run client.go
```

### 2. Verify

#### ✅ Check Kafka

Open [http://localhost:9000](http://localhost:9000) (Kafdrop) and view the `location-updates` topic.

#### ✅ Check Cassandra

```bash
docker exec -it cassandra cqlsh -e "SELECT * FROM location.locations;"
```
