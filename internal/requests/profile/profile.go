package profile

import (
	cerr "SocialServiceAincrad/custom_errors"
	"SocialServiceAincrad/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func ProfileGET(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.Unauthorized.Error()})
		return
	}

	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"message": "Profile Page"})
}

// POST
func ProfilePOST(c *gin.Context) {

}
