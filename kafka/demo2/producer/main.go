package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "mydemo1",
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 100 * time.Millisecond,
	}

	n := 10
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			err := w.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte("Key-B"),
					Value: []byte("5Hello World!"),
				},
				kafka.Message{
					Key:   []byte("Key-B"),
					Value: []byte("5One!"),
				},
				kafka.Message{
					Key:   []byte("Key-B"),
					Value: []byte("5Two!"),
				},
				kafka.Message{
					Key:   []byte("Key-B"),
					Value: []byte("yangjie!"),
				},
			)

			if err != nil {
				log.Fatal("failed to write messages:", err)
			}
		}()
	}

	wg.Wait()

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
