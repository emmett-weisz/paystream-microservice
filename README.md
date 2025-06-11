# Paystream Microservice

A simple Go microservice that exposes a gRPC endpoint to verify payments and publishes messages to a Kafka topic.

---

## Features

- CLI interface powered by [Cobra](https://github.com/spf13/cobra)
- Configuration management with [Viper](https://github.com/spf13/viper)
- gRPC server with payment verification endpoint
- Kafka producer publishes messages on payment verification

---

## Requirements

- Go 1.18+
- Docker & Docker Compose
- Kafka & Zookeeper running (can be done via Docker Compose)
- Protobuf compiler (`protoc`)

---

## Project Highlights

- Written in Go — The entire microservice is implemented in Go.
- CLI with Cobra — cmd/root.go defines a CLI using Cobra to parse arguments and start the server.
- Config management with Viper — Viper is used to load configuration (e.g. Kafka topic name, gRPC port).
- gRPC server with custom proto — Implements a VerifyPayment gRPC endpoint, defined in proto/payment.proto.
- Kafka integration — When the gRPC endpoint is called, the service publishes a structured message to a Kafka topic using kafka-go.
- End-to-end tested — Messages were successfully verified via grpcurl and consumed from Kafka within the container to confirm full functionality.



---
## Setup & Run

1. Clone the repo:

   `git clone https://github.com/yourusername/paystream-microservice.git`
   `cd paystream-microservice`

2. Start Kafka & Zookeeper using Docker Compose:

  `docker-compose up -d` and give it a few seconds to load

3. Create your kafka topic: "verified-payments"

  `docker exec -it paystream-microservice-kafka-1 kafka-topics.sh \
  --create \
  --topic verified-payments \
  --bootstrap-server localhost:9092 \
  --partitions 1 \
  --replication-factor 1`

3. Build and run the gRPC server:

  `go run main.go`

4. Test the gRPC endpoint 

  # Option 1: Using grpcurl

  grpcurl -plaintext -d '{
  "payer_id": "emmettweisz",
  "amount": 100.50,
  "currency": "USD",
  "payment_method": "credit_card"
  }' localhost:50051 payment.PaymentVerifier.VerifyPayment

  # Option 2: Using the go client

  `go run ./examples/client.go`


## Expected Response Using grpcurl

  {
    "status": "success",
    "message": "Payment verified and message published"
  }

## Expected Response Using the the Go Client

  `2025/06/10 20:24:57 Response: status:"success" message:"Payment verified and message published`

## Prove it!

  `docker ps` and grab your containerId

  `docker exec -it  <containerId> /bin/bash`

  `kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic verified-payments --from-beginning`


## Point of Contact

- Created by Emmett Weisz