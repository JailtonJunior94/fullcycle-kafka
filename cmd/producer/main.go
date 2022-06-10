package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	producer := NewKafkaProducer()
	if err := Publish("Mensagem via GO", "teste", producer, nil, deliveryChan); err != nil {
		log.Println(err.Error())
	}

	go DeliveryReport(deliveryChan)
	producer.Flush(2000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "localhost:9093",
		"delivery.timeout.ms": "0",
		"acks":                "1",
		"enable.idempotence":  "false",
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return producer
}

func Publish(message, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	msg := &kafka.Message{
		Value: []byte(message),
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	if err := producer.Produce(msg, deliveryChan); err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("erro ao enviar mensagem")
				return
			}
			fmt.Println("mensagem enviada: ", ev.TopicPartition)
			return
		}
	}
}
