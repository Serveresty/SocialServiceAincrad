package handlers

import (
	signin "SocialServiceAincrad/internal/sign_in"
	signup "SocialServiceAincrad/internal/sign_up"

	"github.com/gin-gonic/gin"
)

func AllRequests(router *gin.Engine) {
	router.GET("/sign-in", signin.SignIn)
	router.GET("/sign-up", signup.SignUp)
}
