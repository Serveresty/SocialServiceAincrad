package handlers

import (
	signin "SocialServiceAincrad/internal/requests/sign_in"
	signup "SocialServiceAincrad/internal/requests/sign_up"

	"github.com/gin-gonic/gin"
)

func AllRequests(router *gin.Engine) {
	router.GET("/sign-in", signin.SignInGET)
	router.POST("/login", signin.SignInPOST)

	router.GET("/sign-up", signup.SignUpGET)
	router.POST("/registration", signup.SignUpPOST)
}
