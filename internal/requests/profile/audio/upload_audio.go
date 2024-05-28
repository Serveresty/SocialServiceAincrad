package audio

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadAudioPOST(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrUnauthorized.Error()})
		return
	}

	claims, err := jwtservice.ParseToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("audio")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := c.PostForm("author")
	title := c.PostForm("title")
	song := models.Audio{
		Name:   title,
		Author: author,
	}

	hash, err := utils.GetFileHash(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var filename string
	var songID int
	for attempt := 0; attempt < 3; attempt++ {
		filename, err = utils.GenerateFilename()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		songID, err = profiledb.SetAudioToUpload(song, hash, filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			if attempt == 2 {
				return
			}
			continue
		}
		break
	}

	err = utils.SaveFileToStorage(c, filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = profiledb.SetAudioToFavorite(claims.Subject, songID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "upload success"})
}
