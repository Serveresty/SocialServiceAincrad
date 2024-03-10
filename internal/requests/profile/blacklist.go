package profile

import (
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlacklistGET(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrUnauthorized.Error()})
		return
	}

	// claims, err := jwtservice.ParseToken(c)
	// if err != nil {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	// 	return
	// }
}

func BlacklistPOST(c *gin.Context) {

}
