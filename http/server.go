package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine  *gin.Engine
	handler Handler
}

type Handler struct {
}

type reg struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func NewServer() Server {
	router := gin.Default()
	handler := Handler{}

	router.GET("/hello", handler.Hello)
	router.POST("/user", handler.CreateUser)
	return Server{
		engine:  router,
		handler: handler,
	}
}

func (H Handler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func (s *Server) Run(puerto string) error {
	return s.engine.Run(puerto)

}

func (H Handler) CreateUser(c *gin.Context) {
	var request reg
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
