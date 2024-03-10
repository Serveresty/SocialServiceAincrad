package signin

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET
func SignInGET(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sign-In Page"})
}

// POST
func SignInPOST(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}

	var user models.AuthUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := profiledb.GetAuthData(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roles, err := profiledb.GetUserRoles(id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error while scanning roles: " + err.Error()})
		return
	}

	token, err := jwtservice.GenerateToken(strconv.Itoa(id), roles, user.StayLoggedIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
