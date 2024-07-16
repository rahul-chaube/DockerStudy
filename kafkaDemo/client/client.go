package client

import (
	"KafkaDemo/model"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	KafkaServer  = "localhost:9092"
	KafkaTopic   = "orders-v1-topic"
	KafkaGroupId = "product-service"
)

func InitClient() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
		"group.id":          KafkaGroupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println("Error occured ******  111 ")
		panic(err)
	}
	defer c.Close()

	topic := KafkaTopic

	c.SubscribeTopics([]string{topic}, nil)
	fmt.Println("Connected ")
	for {
		fmt.Println("Message received #######  ")
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var order model.Order
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				fmt.Printf("11111   Error decoding message: %v\n", err)
				continue
			}

			fmt.Printf("Received Order: %+v\n", order)
		} else {
			fmt.Printf(" 222222  Error: %v\n", err)
		}
	}

}
