package server

import "github.com/gin-gonic/gin"

type Endpoint interface {
	Handler(c *gin.Context)
}
