package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	models.User
	LoginAt  time.Time `json:"loginAt"`
	ExpireAt time.Time `json:"expireAt"`
	Token    string    `json:"token"`
}

func Signup(c *gin.Context) {
	var userBody UserRequestBody

	if c.Bind(&userBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{
		Email:    userBody.Email,
		Password: string(hash),
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success"})
}

func LoginPlaintextPasswordJWT(c *gin.Context) {
	//get the email & password from the request body

	//登录session 有效期 20min
	var sessionTime time.Duration = 20 * time.Minute

	var userBody UserRequestBody

	if c.Bind(&userBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	//look up request email in the database

	var user models.User
	db.DB.First(&user, "email = ? and status = 0", userBody.Email)

	var loginResponse LoginResponse

	// 如果用户不存在，返回错误
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password or email",
		})
		return
	}
	//compare the password in the database with the request password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userBody.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password or email",
		})
		return
	}

	//generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(sessionTime).Unix(),
	})

	secret := viper.GetString("jwt.secret")
	tokenString, err := token.SignedString([]byte(secret))
	//return the JWT token

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	loginResponse.Email = user.Email
	loginResponse.LoginAt = time.Now()
	loginResponse.ExpireAt = time.Now().Add(sessionTime)

	loginResponse.Token = tokenString
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, loginResponse)
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged in",
	})
}

// forget password
func ForgetPassword(c *gin.Context) {
	//get the email from the request body
	var email string
	if c.Bind(&email) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	//look up request email in the database
	var user models.User
	db.DB.First(&user, "user_id = ?", email)
	// 如果用户不存在，返回错误
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}

	//如果用户 status 为 1，返回错误
	if user.Status == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user is inactive",
		})
		return
	}

	//如果用户 status 为 0，则 设置用户 status 为 1，返回成功
	user.Status = 1
	db.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// reset password
func ResetPassword(c *gin.Context) {
	//get the request from the request body
	var request struct {
		UserRequestBody
		NewPassword string `json:"newPassword"`
	}
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	//look up request email in the database
	var user models.User
	db.DB.First(&user, "user_id = ?", request.Email)
	// 如果用户不存在，返回错误
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}

	//如果用户 status 不为 0，返回错误
	if user.Status != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user is inactive",
		})
		return
	}

	//计算密码的 hash，判断是否与数据库中的密码相同
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password",
		})
		return
	}

	//如果相同，计算新密码的 hash，更新数据库中的密码
	hash, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	user.Password = string(hash)
	db.DB.Save(&user)

	//返回成功
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
