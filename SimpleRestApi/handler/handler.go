package handler

import (
	"SimpleRestApi/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type RestInterface interface {
	GetUser(http.ResponseWriter, *http.Request)
	AddUser(http.ResponseWriter, *http.Request)
}

type RestService struct {
}

func RestServiceHandler() *RestService {
	return &RestService{}
}

func (r *RestService) GetUser(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		http.Error(res, "Invalid request method ", http.StatusMethodNotAllowed)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(model.UserData)
	res.Write(resp)

}

func (r *RestService) AddUser(res http.ResponseWriter, req *http.Request) {
	var user model.User

	if req.Method != http.MethodPost {
		http.Error(res, "Invalid request method ", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body", http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(res, "Error reading request body", http.StatusInternalServerError)
		return
	}
	user.Id = uuid.NewString()
	model.UserData = append(model.UserData, user)
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusCreated)

	fmt.Println(res.Header())
	resp, _ := json.Marshal(user)
	res.Write(resp)

}
