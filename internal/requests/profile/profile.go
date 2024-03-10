package profile

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	profilemethods "SocialServiceAincrad/internal/requests/profile/profile_methods"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET
func ProfileGET(c *gin.Context) {
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

	id := c.Param("id")

	var profileData *models.ProfileData
	if claims.Subject == id {
		profileData, err = profilemethods.GetProfileData(id, models.PrivacySettings{})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	} else {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		privacy, err := profiledb.GetPrivacySettings(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		profileData, err = profilemethods.GetProfileData(id, *privacy)
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
