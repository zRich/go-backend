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

func Signup(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{
		Email:    &body.Email,
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
		"message": "success",
	})
}

func LoginPlaintextPasswordJWT(c *gin.Context) {
	//get the email & password from the request body

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to parse body",
		})
		return
	}

	//look up request email in the database

	var user models.User
	db.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password or email",
		})
		return
	}
	//compare the password in the database with the request password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password or email",
		})
		return
	}

	//generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
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

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString},
	)
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged in",
	})
}
