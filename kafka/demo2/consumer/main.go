package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "md15",
		Topic:       "mydemo1",
		MaxWait:     500 * time.Millisecond,
		StartOffset: kafka.FirstOffset,
	})

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				break
			}
			fmt.Printf("topic/partition/offset %v/%v/%v: %s = %s \n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	if err := r.Close(); err != nil {
		log.Fatal(err)
	}
}
