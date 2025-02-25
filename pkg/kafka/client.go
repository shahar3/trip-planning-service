package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Client struct {
	Writer *kafka.Writer
}

func NewKafkaClient(brokers []string, topic string) *Client {
	return &Client{
		Writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (kc *Client) SendMessage(ctx context.Context, key, message []byte) error {
	return kc.Writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: message,
	})
}
