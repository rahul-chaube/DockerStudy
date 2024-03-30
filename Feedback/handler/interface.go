package handler

import "github.com/kataras/iris/v12"

type HandlerInterface interface {
	AddFeedback(iris.Context)
	ListFaadback(iris.Context)
}

type Feedback struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
