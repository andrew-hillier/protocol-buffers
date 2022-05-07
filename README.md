# protocol-buffers

## Step One - Kafka
https://developer.confluent.io/get-started/go/#introduction

```
docker compose up -d

docker compose exec broker \
  kafka-topics --create \
    --topic purchases \
    --bootstrap-server localhost:9092 \
    --replication-factor 1 \
    --partitions 1

go build -o out/producer util.go producer.go

go build -o out/consumer util.go consumer.go
```