package logout

import (
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrUnauthorized.Error()})
		return
	}

	c.Request.Header.Del("Authorization")
	c.JSON(http.StatusOK, gin.H{"message": "Logout success"})
}
