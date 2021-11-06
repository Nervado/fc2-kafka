package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	deliveryChannel := make(chan kafka.Event)

	producer := NewKafkaProducer()

	publish("ola mundo", "teste", producer, nil, deliveryChannel)

	e := <-deliveryChannel
	msg := e.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar ")
	} else {
		fmt.Println("Messagem enviada: ", msg)
	}

	producer.Flush(1000) // wait publisher
}

func NewKafkaProducer() *kafka.Producer {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "fc2-kafka_kafka_1:9092",
	}

	p, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}

	return p
}

func publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(message, deliveryChannel)

	if err != nil {
		log.Println(err.Error())
	}

	return nil

}

func DeliveryReport(deliveryChannel chan kafka.Event) {}
