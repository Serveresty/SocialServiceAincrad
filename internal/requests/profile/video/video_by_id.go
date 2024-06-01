package video

import (
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetVideoByID(c *gin.Context) {
	//id := c.Param("id")
	vid := c.Param("vid")

	video, err := profiledb.GetVideoById(vid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}

	filePath := "../../storages/video_storage/" + video.Filename + ".mp4"
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video file not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	dataSize := strconv.FormatInt(info.Size(), 10)

	err = profiledb.AddViewToVideo(vid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "video/mp4")
	c.Header("Content-Length", dataSize)
	c.Header("Content-Disposition", "inline; filename="+video.Filename+".mp4")

	c.File(filePath)
}
