package server

import (
	"KafkaDemo/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

const (
	KafkaServer = "localhost:9092"
	KafkaTopic  = "orders-v1-topic"
)

func InitServer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
	})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := KafkaTopic

	for {
		order := model.Order{
			Id:      uuid.NewString(),
			Name:    "Marriedgold",
			UserId:  "abc",
			Ammount: 5000,
		}

		value, err := json.Marshal(order)

		if err != nil {
			panic(err)
		}
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          value,
		}, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Message Send  ^^^^^^^^^^^  ")
		time.Sleep(time.Second)

	}

}
