package signup

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	utilsdb "SocialServiceAincrad/internal/database/utils_db"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"net/http"
	"os"

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

	// err1 := os.Mkdir("../../storages/audio_storage/"+user.Username, 0755)
	// if err1 != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating audio storage: " + err1.Error()})
	// 	return
	// }
	err2 := os.Mkdir("../../storages/photo_storage/"+user.Username, 0755)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating photo storage: " + err2.Error()})
		return
	}
	err3 := os.Mkdir("../../storages/video_storage/"+user.Username, 0755)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating video storage: " + err3.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
}
