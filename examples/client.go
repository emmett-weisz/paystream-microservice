package main

import (
	"context"
	"log"
	"time"

	pb "github.com/emmett-weisz/paystream-microservice/proto/paymentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPaymentVerifierClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.VerifyPayment(ctx, &pb.PaymentRequest{
		PayerId:       "emmettweisz",
		Amount:        100.50,
		Currency:      "USD",
		PaymentMethod: "credit_card",
	})
	if err != nil {
		log.Fatalf("error calling VerifyPayment: %v", err)
	}

	log.Printf("Response: %+v", resp)
}
