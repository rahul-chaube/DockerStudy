package main

import (
	"KafkaDemo/server"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start Server ******* ")

	go server.InitServer()

	// fmt.Println("Starting client ")

	// go client.InitClient()

	// fmt.Println("Print log *********  ")

	time.Sleep(time.Minute)

}
