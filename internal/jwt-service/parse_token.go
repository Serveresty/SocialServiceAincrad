package jwtservice

import (
	"SocialServiceAincrad/configs"
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/models"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ParseToken(c *gin.Context) (*models.JWTClaims, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return nil, cerr.ErrUnauthorized
	}
	tokenString := strings.Split(token, " ")[1]

	secretKey := []byte(configs.GetEnv("SECRET_KEY"))

	currentToken, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := currentToken.Claims.(*models.JWTClaims)
	if !ok {
		return nil, cerr.ErrorClaims
	}

	return claims, nil
}
