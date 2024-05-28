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
		api.GET("sign-in", signin.SignInGET) //Did it
		api.POST("login", signin.SignInPOST) //Did it

		//Регистрация
		api.GET("sign-up", signup.SignUpGET)        //Did it
		api.POST("registration", signup.SignUpPOST) //Did it

		//Выход
		api.GET("logout", logout.Logout) //Did it

		//Профиль
		api.GET(":id", profile.ProfileGET)   //Did it
		api.POST(":id", profile.ProfilePOST) // ?

		//Список друзей
		api.GET("friends", profile.FriendsGET) //Did it

		//Музыка
		api.GET("audio", audio.AudioGET)         //Did it not full
		api.GET("audio/:id", audio.GetAudioById) //Did it
		api.POST("audio/upload", audio.UploadAudioPOST)
		api.POST("audio/to-favorite", audio.AudioToFavoritePOST)
		api.POST("audio/delete-favorite", audio.AudioDeleteFavoritePOST)

		//Видео
		api.GET("video", video.VideoGET)
		api.GET("video/:id", video.VideoCurrentUserGET)
		api.POST("video/upload")
		api.POST("video/to-favorite")
		api.POST("video/delete-favorite")

		//Чаты
		api.GET("messages", messages.ChatGET) //Did it

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
