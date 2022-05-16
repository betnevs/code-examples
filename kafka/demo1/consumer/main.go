package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "md3",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}

	topic := "mydemo1"

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			ev, err := consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("read error: %v", err)
				continue
			}
			fmt.Printf("comsumed from topic: %s: key = %-10s value = %s\n", ev.TopicPartition,
				string(ev.Key), string(ev.Value))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		fmt.Printf("Caught signal: %v", sig)
	}

	err = consumer.Close()
	if err != nil {
		log.Fatal(err)
	}

}
