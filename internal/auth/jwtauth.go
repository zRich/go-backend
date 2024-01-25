package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
)

func JWTAuth(c *gin.Context) {
	token, err := GetTokenFromCookie(c)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		db.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func GetTokenFromCookie(c *gin.Context) (*jwt.Token, error) {
	//get the token from the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		return nil, err
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpeted signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})

	return token, nil
}

func GetUserFromToken(token *jwt.Token) (*models.User, error) {
	var user models.User

	//validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, fmt.Errorf("token expired")
		}
		db.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			return nil, fmt.Errorf("user not found")
		}
	} else {
		return nil, fmt.Errorf("invalid token")
	}

	return &user, nil
}
