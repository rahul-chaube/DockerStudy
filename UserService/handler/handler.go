package handler

import (
	"UserService/model"
	"UserService/service"
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	service *service.UserService
}

func InitUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: &service,
	}
}
func (h *UserHandler) AddUser(ctx iris.Context) {
	start := time.Now()
	// h.service.Log.Debug("Start Wroking like pro")

	var user model.User

	err := ctx.ReadBody(&user)
	if err != nil {
		ctx.StopWithPlainError(http.StatusBadRequest, iris.NewProblem().Title("User Creation is failed ").DetailErr(err))
		return
	}
	user.LastUpdate = time.Now()
	user.AddedAt = time.Now()

	h.service.AddUser(user)
	h.service.Log.WithFields(logrus.Fields{"Name": user.Name, "Age": user.Age, "time": time.Now(), "timeTaken": time.Since(start)}).Debug("User add done ")
	ctx.StatusCode(http.StatusCreated)
	ctx.WriteString("User is created successfully")

}
