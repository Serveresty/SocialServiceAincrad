package utils

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func SaveFileToStorage(c *gin.Context, filename string, file *multipart.FileHeader) error {
	err := c.SaveUploadedFile(file, "../../storages/audio_storage/"+filename+".mp3")
	if err != nil {
		return err
	}
	return nil
}
