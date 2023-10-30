package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/auth"
)

type ServerConfig struct {
	Port int
}

type Server interface {
	Start() error
}

type HttpServer struct {
	Config *ServerConfig
}

func (s *HttpServer) Start() error {
	r := gin.Default()
	r.POST("/signup", auth.Signup)
	r.POST("/login", auth.Login)
	r.GET("/validate", auth.RequireAuth, auth.Validate)
	r.Use(cors.Default())
	return r.Run(fmt.Sprintf(":%d", s.Config.Port))
}

func NewServer(config *ServerConfig) Server {
	httpserver := &HttpServer{
		Config: config,
	}
	return httpserver
}
