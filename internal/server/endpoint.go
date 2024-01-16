package server

import "github.com/gin-gonic/gin"

type Endpoint interface {
	Method() string
	Path() string
	LoginVerify() bool
	Handler(ctx *gin.Context)
}
