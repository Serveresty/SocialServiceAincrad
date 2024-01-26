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
		api.POST(":id", profile.ProfilePOST) // ?

		settings := api.Group("settings")
		{
			settings.GET("", profile.ProfileSettingsGeneralGET)
			settings.POST("save-general", profile.ProfileSettingsGeneralPOST)

			settings.GET("privacy", profile.ProfileSettingsPrivacyGET)
			settings.POST("save-privacy", profile.ProfileSettingsPrivacyPOST)

			settings.GET("blacklist", profile.BlacklistGET)
			settings.POST("save-blacklist", profile.BlacklistPOST)
		}
	}
}
