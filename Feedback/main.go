package main

import (
	"feedback/handler"

	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.New()

	handler := handler.NewFeedbackHandler()

	feedback := app.Party("/feedback")
	{
		feedback.Post("", handler.AddFeedback)
	}
	app.Listen(":8081")
}
