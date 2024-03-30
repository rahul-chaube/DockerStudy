package handler

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/kataras/iris/v12"
)

type handler struct {
}

func NewFeedbackHandler() *handler {

	return &handler{}
}

func (h *handler) AddFeedback(ctx iris.Context) {
	var feedback Feedback

	err := ctx.ReadBody(&feedback)
	if err != nil {
		ctx.StopWithPlainError(http.StatusBadRequest, iris.NewProblem().Title("Failed To add Feedback ").DetailErr(err))
		return
	}
	err = os.WriteFile("feedback/"+strings.ToLower(feedback.Title)+".txt", []byte(feedback.Description), fs.FileMode(os.O_CREATE))
	if err != nil {
		fmt.Println("Error is ", err)
		ctx.StopWithPlainError(http.StatusBadRequest, iris.NewProblem().Title("Failed To add Feedback ").DetailErr(err))
		return
	}

	ctx.StatusCode(http.StatusCreated)
	ctx.WriteString("File is created")
}

func (h *handler) ListFaadback(iris.Context) {

}
