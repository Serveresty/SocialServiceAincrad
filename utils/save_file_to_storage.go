package utils

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func SaveFileToStorage(c *gin.Context, filename string, storageType string, fileType string, file *multipart.FileHeader) error {
	err := c.SaveUploadedFile(file, "../../storages/"+storageType+"/"+filename+fileType)
	if err != nil {
		return err
	}
	return nil
}
