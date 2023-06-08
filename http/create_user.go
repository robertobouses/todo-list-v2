package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerTask struct {
	http.Handler
}
type reg struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func (H HandlerTask) CreateUserRequest(c *gin.Context) {
	var request reg
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
