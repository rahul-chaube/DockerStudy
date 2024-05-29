package handler

import (
	"MovieService/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MoviesInterface interface {
	GetMoviesList(iris.Context)
	GetPeople(ctx iris.Context)
}

type MovieService struct {
	Log       *logrus.Logger
	mngClient *mongo.Client
}

func NewHandler(log *logrus.Logger, mngClient *mongo.Client) MovieService {
	return MovieService{
		Log:       log,
		mngClient: mngClient,
	}
}

func (m *MovieService) GetMoviesList(ctx iris.Context) {
	data, err := http.Get("https://swapi.dev/api/films")
	if err != nil {
		fmt.Println("Error occured ", err.Error())
		m.Log.Error("Error is occured ", err.Error())
		ctx.StopWithError(http.StatusBadRequest, iris.NewProblem().Title("Failed to retrive movies list ").DetailErr(err))
		return
	}

	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	ctx.Header("Content-Type", "application/json")
	ctx.StatusCode(http.StatusOK)
	ctx.Write(body)
}

func (m *MovieService) GetPeople(ctx iris.Context) {
	id := ctx.Params().Get("id")

	fmt.Println("People id ", id)

	data, err := http.Get("https://swapi.dev/api/people/" + id)
	if err != nil {
		fmt.Println("Error occured ", err.Error())
		m.Log.Error("Error is occured ", err.Error())
		ctx.StopWithError(http.StatusBadRequest, iris.NewProblem().Title("Failed to retrive movies list ").DetailErr(err))
		return
	}

	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	ctx.Header("Content-Type", "application/json")
	ctx.StatusCode(http.StatusOK)
	ctx.Write(body)
}

func (m *MovieService) AddFavouratePeople(ctx iris.Context) {

	var favourate model.Favourate

	err := ctx.ReadBody(&favourate)
	if err != nil {
		ctx.StopWithError(http.StatusBadRequest, iris.NewProblem().Title("Failed to accept request").DetailErr(err))
		return
	}
	favourate.Id = uuid.NewString()
	_, err = m.mngClient.Database("MovieRepository").Collection(model.UserTable).InsertOne(ctx, favourate)
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, iris.NewProblem().Title("Failed to insert in databse ").DetailErr(err))
		return
	}
	m.Log.Info(favourate)

	ctx.Header("Content-Type", "application/json")
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"Message": "User Added successfully"})
}

func (m *MovieService) GetFavouratePeopleList(ctx iris.Context) {
	var favouriteList = []model.Favourate{}

	cursor, err := m.mngClient.Database("MovieRepository").Collection(model.UserTable).Find(ctx, bson.M{})
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, iris.NewProblem().Title("Failed to insert in databse ").DetailErr(err))
		return
	}
	for cursor.Next(ctx) {
		fmt.Println("Itterate ")
		var fav model.Favourate
		if err := cursor.Decode(&fav); err != nil {
			fmt.Println("Error occred ", err)
			break
		}
		fmt.Println("Found favourate ", fav)
		favouriteList = append(favouriteList, fav)
	}

	// var result []bson.M
	// if err = cursor.All(ctx, &result); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, data := range result {

	// 	fmt.Println(" Data center issue ", data)
	// }

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(favouriteList)
}
