package handlers

import (
	signin "SocialServiceAincrad/internal/sign_in"
	signup "SocialServiceAincrad/internal/sign_up"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func AllRequests(router *gin.Engine, dbConn *pgx.Conn) {
	router.GET("/sign-in", signin.SignIn)
	router.GET("/sign-up", signup.SignUp)
}
