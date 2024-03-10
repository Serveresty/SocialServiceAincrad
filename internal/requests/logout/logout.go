package logout

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.Request.Header.Del("Authorization")
	c.JSON(http.StatusOK, gin.H{"message": "Logout success"})
}
