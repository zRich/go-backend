package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/auth"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
	"github.com/zRich/go-backend/internal/server"
)

type GetUserAdminEndpoint struct{}

// Method returns the HTTP method for this endpoint.
func (e *GetUserAdminEndpoint) Method() string {
	return "GET"
}

// Path returns the HTTP path for this endpoint.
func (e *GetUserAdminEndpoint) Path() string {
	return "useradmin"
}

// LoginVerify returns whether this endpoint requires login verification.
func (e *GetUserAdminEndpoint) LoginVerify() bool {
	return true
}

// Handler returns the handler function for this endpoint.
func (e GetUserAdminEndpoint) Handler(ctx *gin.Context) {
	var user models.User
	var response server.HttpResonpose

	_end, _ := strconv.Atoi(ctx.DefaultQuery("_end", "10"))
	_start, _ := strconv.Atoi(ctx.DefaultQuery("_start", "1"))

	//get token from header
	token, err := auth.GetTokenFromCookie(ctx)

	if err != nil {
		//return unauthorized
		response = server.WrapResponse(http.StatusUnauthorized, err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	//get user from token
	admin, err := auth.GetUserFromToken(token)
	if err != nil {
		//return unauthorized
		response = server.WrapResponse(http.StatusUnauthorized, err.Error(), nil)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	// is admin? Role = 1, Status = 0
	if admin.Role != 1 || admin.Status != 0 {
		response = server.WrapResponse(http.StatusUnauthorized, "not admin", nil)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	if err := db.DB.Limit(_end - _start).Offset(_start).Find(&user).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
