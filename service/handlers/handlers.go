package handlers

import (
	"SocialServiceAincrad/internal/requests/logout"
	"SocialServiceAincrad/internal/requests/profile"
	"SocialServiceAincrad/internal/requests/profile/audio"
	"SocialServiceAincrad/internal/requests/profile/messages"
	"SocialServiceAincrad/internal/requests/profile/video"
	signin "SocialServiceAincrad/internal/requests/sign_in"
	signup "SocialServiceAincrad/internal/requests/sign_up"

	"github.com/gin-gonic/gin"
)

func AllRequests(router *gin.Engine) {
	api := router.Group("")
	{
		//Авторизация
		api.GET("sign-in", signin.SignInGET)
		api.POST("login", signin.SignInPOST)

		//Регистрация
		api.GET("sign-up", signup.SignUpGET)
		api.POST("registration", signup.SignUpPOST)

		//Выход
		api.POST("logout", logout.Logout)

		//Профиль
		api.GET(":id", profile.ProfileGET)
		api.POST(":id", profile.ProfilePOST) // ?

		//Список друзей
		api.GET("friends", profile.FriendsGET)

		//Музыка
		api.GET("audio", audio.AudioGET)
		api.GET("audio/:id", audio.GetAudioById)

		//Видео
		api.GET("video", video.VideoGET)
		api.GET("video/:id", video.VideoCurrentUserGET)

		//Чаты
		api.GET("messages", messages.ChatGET)

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
