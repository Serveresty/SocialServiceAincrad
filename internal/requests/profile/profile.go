package profile

import (
	cerr "SocialServiceAincrad/custom_errors"
	profileactions "SocialServiceAincrad/internal/database/profile_actions"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
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

	claims, err := jwtservice.ParseToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	var profileData *models.ProfileData
	if claims.Id == id {
		profileData, err = profileactions.GetProfileData(claims.Id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	} else {
		// Check private settings then get page!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		profileData, err = profileactions.GetProfileData(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": profileData})
}

// POST
func ProfilePOST(c *gin.Context) {

}
