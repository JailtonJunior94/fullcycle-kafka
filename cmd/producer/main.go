package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer := NewKafkaProducer()
	if err := Publish("mensagem via GO", "teste", producer, nil); err != nil {
		log.Println(err.Error())
	}
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return producer
}

func Publish(message, topic string, producer *kafka.Producer, key []byte) error {
	msg := &kafka.Message{
		Value: []byte(message),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	if err := producer.Produce(msg, nil); err != nil {
		return err
	}
	return nil
}
