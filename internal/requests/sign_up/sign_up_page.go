package signup

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	utilsdb "SocialServiceAincrad/internal/database/utils_db"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func SignUpGET(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sign-Up Page"})
}

// POST
func SignUpPOST(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok := utilsdb.IsUserRegistered(&user)
	if ok {
		c.JSON(http.StatusConflict, gin.H{"error": "user already registered"})
		return
	}

	hashPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while hashing password: " + err.Error()})
		return
	}

	user.Password = hashPwd

	err = profiledb.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while creating a new user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
}
