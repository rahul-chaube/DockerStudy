package main

import (
	"MovieService/handler"
	"MovieService/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// file, err := os.OpenFile(config.MovieService.LogDir+config.MovieService.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	logrus.Fatal("Failed to open log file:", err)
	// }
	// defer file.Close()

	// log.SetOutput(file)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	client, err := mongoInit(config)
	if err != nil {
		log.Info("Mongo Url: ", config.MovieService.MongoURL)
		log.Fatal(err, config.MovieService.MongoURL)
	}
	log.Info("Mongodb Client connected successfully ")
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	movieService := handler.NewHandler(log, client)

	movie := app.Party("movies")
	{
		movie.Get("", movieService.GetMoviesList)
	}

	people := app.Party("people")
	{
		people.Get("/{id}", movieService.GetPeople)
	}

	favourate := app.Party("favourate")
	{
		favourate.Get("", movieService.GetFavouratePeopleList)
		favourate.Post("", movieService.AddFavouratePeople)
	}
	app.Listen(fmt.Sprintf(":%d", config.MovieService.PortNumber))

}

func mongoInit(config model.Config) (*mongo.Client, error) {
	ctx := context.Background()
	// Set client options
	clientOptions := options.Client().ApplyURI(config.MovieService.MongoURL)
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err

	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// coll := client.Database("UserDatabase").Collection("User")
	// cursor, err := coll.Find(ctx, bson.D{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(ctx)

	// var result []bson.M
	// if err = cursor.All(ctx, &result); err != nil {
	// 	log.Fatal(err)
	// }

	// var FavourateList []model.Favourate

	// for _, data := range result {

	// 	fmt.Println(" Data center issue ", data)
	// }

	return client, nil
}
