package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type PaymentMessage struct {
	PayerID       string  `json:"payer_id"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	PaymentMethod string  `json:"payment_method"`
}

func SendPaymentMessage(writer *kafka.Writer, msg PaymentMessage) error {
	value, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(msg.PayerID),
			Value: value,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %v", err)
		return err
	}

	log.Println("Message published to Kafka")
	return nil
}
