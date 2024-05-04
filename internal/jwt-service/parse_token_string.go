package jwtservice

import (
	"SocialServiceAincrad/configs"
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/models"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ParseTokenString(token string) (*models.JWTClaims, error) {
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
