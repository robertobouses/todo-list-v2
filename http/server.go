package http

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
}

func NewServer() Server {
	return Server{
		engine: gin.Default(),
	}
}

func (s *Server) Run(puerto string) error {
	return s.engine.Run()
}
