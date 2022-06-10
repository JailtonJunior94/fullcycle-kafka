package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"client.id":         "goapp0-consumer",
		"group.id":          "goapp0-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("Erro consumer", err.Error())
	}

	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}
