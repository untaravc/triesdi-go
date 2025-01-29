package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

type ClaimData struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func GenerateToken(id string, email string, phone string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &ClaimData{
		Id:    id,
		Email: email,
		Phone: phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*ClaimData, error) {
	claims := &ClaimData{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func GetAuth(ctx *gin.Context) ClaimData {
	data_auth, _ := ctx.Get("auth_user")
	return data_auth.(ClaimData)
}
