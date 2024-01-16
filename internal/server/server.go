package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/auth"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/log"
)

type Server interface {
	Start() error
	AddEndpoint(endpoint Endpoint)
}

type HttpResonpose struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HttpServer struct {
	DB     db.Database
	Config *RestConfig
	engine *gin.Engine
}

// 返回 server 的engine
func (s *HttpServer) Engine() *gin.Engine {
	return s.engine
}

func (s *HttpServer) Start() error {
	r := gin.Default()

	log.Log.Info("server start")
	r.Use(cors.Default())
	return r.Run(fmt.Sprintf(":%d", s.Config.Port))
}

func (s *HttpServer) AddEndpoint(endpoint Endpoint) {
	//todo 根据配置文件判断是否需要登录验证和验证方式
	handlers := []gin.HandlerFunc{auth.JWTAuth, endpoint.Handler}
	if !endpoint.LoginVerify() {
		handlers = handlers[1:]
	}
	s.engine.Handle(endpoint.Method(), endpoint.Path(), handlers...)

}

func NewServer(config *RestConfig, DB db.Database) Server {
	restServer := &HttpServer{
		Config: config,
		DB:     DB,
	}
	return restServer
}

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers",
				"Authorization, Content-Length, X-CSRF-Token, Token,session, Content-Type")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers",
				"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()
		c.Next()
	}
}
