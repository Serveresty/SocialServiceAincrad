package jwtservice

import (
	"SocialServiceAincrad/configs"
	"SocialServiceAincrad/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id int, roles []string, loggedIn bool) (string, error) {
	var claims *models.JWTClaims
	if loggedIn {
		claims = &models.JWTClaims{
			Role: roles,
			StandardClaims: jwt.StandardClaims{
				Subject: strconv.Itoa(id),
			},
		}
	} else {
		claims = &models.JWTClaims{
			Role: roles,
			StandardClaims: jwt.StandardClaims{
				Subject:   strconv.Itoa(id),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.GetEnv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenString, nil
}
