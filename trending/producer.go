package trending

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Proschema struct {
	Userid string `json:"Id"`
	Views  int    `json:"Views"`
	UpDate string `json:"UpDate"`
	Days   int    `json:"Days"`
}

func mainPro(jbyte []byte) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "top1"

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(jbyte),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
