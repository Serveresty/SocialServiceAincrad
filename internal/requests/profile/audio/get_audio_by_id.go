package audio

import (
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAudioById(c *gin.Context) {
	id := c.Param("id")

	audio, err := profiledb.GetAudioById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}

	filePath := "../../storages/audio_storage/" + audio.Filename + ".mp3"
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Audio file not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	dataSize := strconv.FormatInt(info.Size(), 10)

	c.Header("Content-Type", "audio/mp3")
	c.Header("Content-Length", dataSize)
	c.Header("Content-Disposition", "inline; filename="+audio.Filename+".mp3")

	c.File(filePath)
}
