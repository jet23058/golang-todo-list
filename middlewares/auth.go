package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	model "todo-list/example/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Id int64 `integer:"id"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var SecretKey = []byte("jet test example")

func GenToken(user model.User) (string, error) {
	c := Auth{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "Jet",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(SecretKey)
}

func ParseToken(tokenString string) (*Auth, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Auth{}, func(token *jwt.Token) (i interface{}, err error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Auth); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "auth is empty.",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "auth error.",
			})
			c.Abort()
			return
		}

		user, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid token.",
			})
			c.Abort()
			return
		}

		c.Set("userId", user.Id)
		c.Next()
	}
}
