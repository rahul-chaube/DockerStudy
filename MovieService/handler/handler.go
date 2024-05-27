package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type MoviesInterface interface {
	GetMoviesList(iris.Context)
	GetPeople(ctx iris.Context)
}

type MovieService struct {
	Log *logrus.Logger
}

func NewHandler() MovieService {
	return MovieService{}
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
