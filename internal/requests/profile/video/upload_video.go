package video

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadVideoPOST(c *gin.Context) {
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

	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	video := models.Video{
		AuthorID:    claims.Subject,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}

	var filename string
	var videoID int
	for attempt := 0; attempt < 3; attempt++ {
		filename, err = utils.GenerateFilename()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		videoID, err = profiledb.SetVideoToUpload(video, filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			if attempt == 2 {
				return
			}
			continue
		}
		break
	}

	storageType := "video_storage"
	fileType := ".mp4"
	err = utils.SaveFileToStorage(c, filename, storageType, fileType, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = utils.CreatePreview("..\\..\\storages\\video_storage\\" + filename + fileType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create preview"})
		return
	}
	err = profiledb.SetPreviewToVideo(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "preview didn't save"})
		return
	}

	err = profiledb.SetVideoToFavorite(claims.Subject, videoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "upload success"})
}
