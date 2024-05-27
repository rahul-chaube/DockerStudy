package main

import (
	"UserService/handler"
	"UserService/model"
	"UserService/service"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func main() {
	app := iris.New()
	// var file, err = os.OpenFile("user.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Println("Could Not Open Log File : " + err.Error())
	// }

	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		return
	}
	var config model.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Println(err)
		return
	}
	log := logrus.New()
	// Create the log file
	file, err := os.OpenFile(config.UserService.LogDir+config.UserService.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	userService := service.InitUserService(log)

	userHandler := handler.InitUserHandler(*userService)

	user := app.Party("/user")
	{
		user.Post("", userHandler.AddUser)
		user.Get("/list", userHandler.ListUser)
	}

	app.Listen(fmt.Sprintf(":%d", config.UserService.PortNumber))

}
