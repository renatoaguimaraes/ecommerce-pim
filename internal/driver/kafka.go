package driver

import (
	"context"
	"fmt"

	"github.com/renatoaguimaraes/ecommerce-pim/internal/util"
	"github.com/segmentio/kafka-go"
)

//Consumer kafka messages
func Consumer(topic string, cgroup string, messages chan []byte) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{util.GetEnv("KAFKA_URL", "localhost:9094")},
		GroupID:  cgroup,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		messages <- m.Value
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	r.Close()
}

var writer *kafka.Writer

// ConnectWriter writer
func ConnectWriter(topic string) {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{util.GetEnv("KAFKA_URL", "localhost:9094")},
		Topic:    topic,
		Async:    false,
		Balancer: &kafka.LeastBytes{},
	})
}

// Send messages
func Send(message kafka.Message) error {
	return writer.WriteMessages(context.Background(), message)
}

//CloseWriter close
func CloseWriter() {
	writer.Close()
}
