package handlers

import (
	"SocialServiceAincrad/internal/requests/profile"
	signin "SocialServiceAincrad/internal/requests/sign_in"
	signup "SocialServiceAincrad/internal/requests/sign_up"

	"github.com/gin-gonic/gin"
)

func AllRequests(router *gin.Engine) {
	api := router.Group("")
	{
		api.GET("sign-in", signin.SignInGET)
		api.POST("login", signin.SignInPOST)

		api.GET("sign-up", signup.SignUpGET)
		api.POST("registration", signup.SignUpPOST)

		api.GET(":id", profile.ProfileGET)
		api.POST(":id", profile.ProfilePOST)
	}
}
