package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
	"golang.org/x/crypto/bcrypt"
)

type SetUserPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}
type SetUserPasswordEndpoint struct {
}

func (e *SetUserPasswordEndpoint) Method() string {
	return "POST"
}

func (e *SetUserPasswordEndpoint) Path() string {
	return "/api/user/password"
}

func (e *SetUserPasswordEndpoint) LoginVerify() bool {
	return true
}

func (e SetUserPasswordEndpoint) Handler(ctx *gin.Context) {
	var request SetUserPasswordRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	token, err := GetTokenFromCookie(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//根据 token 获取用户信息
	admin, err := GetUserFromToken(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//记录用户修改密码的日志
	logger.Info("admin %s set user %s password", admin.Email, request.Email)

	//根据 request.Email 获取用户信息
	var user models.User
	db.DB.First(&user, "email = ?", request.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})
		return
	}

	//计算 request.NewPassword 的 hash, 设置用户密码
	hash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	user.Password = string(hash)
	db.DB.Save(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// active user
type ActiveUserEndpoint struct {
}

func (e *ActiveUserEndpoint) Method() string {
	return "POST"
}

func (e *ActiveUserEndpoint) Path() string {
	return "/api/user/active"
}

func (e *ActiveUserEndpoint) LoginVerify() bool {
	return true
}

func (e ActiveUserEndpoint) Handler(ctx *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	token, err := GetTokenFromCookie(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//根据 token 获取用户信息
	admin, err := GetUserFromToken(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//记录用户修改密码的日志
	logger.Info("admin %s active user %s", admin.Email, request.Email)

	//根据 request.Email 获取用户信息
	var user models.User
	db.DB.First(&user, "email = ? and status <> 0", request.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found or already active",
		})
		return
	}

	//设置用户状态为 0
	user.Status = 0
	db.DB.Save(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// list all inactive users
type ListInactiveUsersEndpoint struct {
}

func (e *ListInactiveUsersEndpoint) Method() string {
	return "GET"
}

func (e *ListInactiveUsersEndpoint) Path() string {
	return "/api/users/inactive"
}

func (e *ListInactiveUsersEndpoint) LoginVerify() bool {
	return true
}

func (e ListInactiveUsersEndpoint) Handler(ctx *gin.Context) {
	var users []models.User
	db.DB.Find(&users, "status <> 0")

	ctx.JSON(http.StatusOK, users)
}
