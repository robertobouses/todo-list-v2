package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine  *gin.Engine
	handler Handler
}

func NewServer() Server {
	router := gin.Default()
	handler := Handler{}

	router.GET("/hello", handler.Hello)
	router.POST("/user", handler.CreateUserRequest)
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
