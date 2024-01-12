package signup

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func SignUpGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Sign-Up Page"})
}

// POST
func SignUpPOST(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok := database.IsUserRegistered(&user)
	if ok {
		c.JSON(http.StatusConflict, gin.H{"error": "user already registered"})
		return
	}
}
