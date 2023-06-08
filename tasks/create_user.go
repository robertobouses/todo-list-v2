package tasks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserParams struct {
	email    string
	password string
}

type reg struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

type HandlerTask struct {
	http.Handler
}

func (t taskAppService) CreateUser(params CreateUserParams) (User, error) {
	if params.email == "" || params.password == "" {
		return User{}, fmt.Errorf("email y password no pueden estar vacíos")
	}

	return User{
		email:    params.email,
		password: params.password,
	}, nil
}

func (H HandlerTask) CreateUserRequest(c *gin.Context) {
	var request reg
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
