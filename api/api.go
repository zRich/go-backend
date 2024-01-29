package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/auth"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/log"
	"github.com/zRich/go-backend/internal/server"
)

type RestServerConfig struct {
	Address string
	Port    int
	Prefix  string
	DB      db.Database
}

func (c *RestServerConfig) GetAddress() string {
	return c.Address
}

func (c *RestServerConfig) GetPort() int {
	return c.Port
}

func (c *RestServerConfig) GetPrefix() string {
	return c.Prefix
}

type RestServer struct {
	DB        db.Database
	Config    server.ServerConfig
	engine    *gin.Engine
	Endpoints []server.Endpoint
}

var logger = log.InitLogger()

// 返回 server 的engine
func (s *RestServer) Engine() *gin.Engine {
	return s.engine
}

func (s *RestServer) Start() error {
	r := s.engine

	r.Use(server.Cors())

	// route group version 1
	v1 := r.Group(s.Config.GetPrefix())
	authRoute := v1.Group("/auth")
	{
		authRoute.POST("/auth/signup", auth.Signup)
		authRoute.POST("/auth/login", auth.LoginPlaintextPasswordJWT)
		authRoute.POST("/auth/forgetPassword", auth.ForgetPassword)
		//reset password
		authRoute.POST("/auth/resetPassword", auth.ResetPassword)
	}

	logger.Info("server start")

	s.initServer()

	endpoints := s.GetEndpoints()

	for _, endpoint := range endpoints {
		s.AddEndpoint(endpoint)
	}

	return r.Run(fmt.Sprintf(":%d", s.Config.GetPort()))
}

func (s *RestServer) AddEndpoint(endpoint server.Endpoint) {
	//todo 根据配置文件判断是否需要登录验证和验证方式
	handlers := []gin.HandlerFunc{auth.JWTAuth, endpoint.Handler}
	if !endpoint.LoginVerify() {
		handlers = handlers[1:]
	}
	s.engine.Handle(endpoint.Method(), fmt.Sprintf("%s/%s", s.Config.GetPrefix(), endpoint.Path()), handlers...)
}

func (s *RestServer) GetEndpoints() []server.Endpoint {
	return s.Endpoints
}

func (s *RestServer) initServer() {
	// course endpoints
	s.Endpoints = append(s.Endpoints, &GetCoursesEndpoint{})
	s.Endpoints = append(s.Endpoints, &CreateCourseEndpoint{})
	s.Endpoints = append(s.Endpoints, &GetCourseByNameEndpoint{})

	// student endpoints
	s.Endpoints = append(s.Endpoints, &GetStudentsEndpoint{})
	s.Endpoints = append(s.Endpoints, &CreateStudentEndpoint{})
	s.Endpoints = append(s.Endpoints, &UpdateStudentByStudentNoEndpoint{})
	s.Endpoints = append(s.Endpoints, &DeleteStudentEndpoint{})

	// task endpoints
	s.Endpoints = append(s.Endpoints, &GetTasksEndpoint{})
	s.Endpoints = append(s.Endpoints, &CreateTaskEndpoint{})
	s.Endpoints = append(s.Endpoints, &UpdateTaskEndpoint{})
	s.Endpoints = append(s.Endpoints, &DeleteTaskEndpoint{})

}

func NewRestServer(config RestServerConfig) server.Server {
	restServer := &RestServer{
		Config: &config,
		DB:     config.DB,
		engine: gin.Default(),
	}
	return restServer
}
