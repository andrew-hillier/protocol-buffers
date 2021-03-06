# protocol-buffers

## Step One - Kafka
https://developer.confluent.io/get-started/go/#introduction

### Dependencies
```
docker compose up -d

docker compose exec kafka \
  kafka-topics --create \
    --topic purchases \
    --bootstrap-server localhost:9092 \
    --replication-factor 1 \
    --partitions 1
```

### Build
```
go build -o out/producer util.go producer.go

go build -o out/consumer util.go consumer.go
```

### Run
```
./out/producer getting-started.properties

./out/consumer getting-started.properties
```

## Step Two - protobuf rw file (non generated)
https://golangbyexample.com/protocol-buffers-go/
https://dev.to/techschoolguru/how-to-define-a-protobuf-message-and-generate-go-code-4g4e

### Build
```
go build -o out/file ioetest/cpu.pb.go iotest/file.go 
```

### Run
```
./out/file
```

## Step Three - generate *.pb.go files
https://developers.google.com/protocol-buffers/docs/gotutorial

```
protoc --go_out=. addressbook.proto
```