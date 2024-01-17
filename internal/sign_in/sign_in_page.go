package signin

import (
	"SocialServiceAincrad/internal/database"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func SignInGET(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sign-In Page"})
}

// POST
func SignInPOST(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	var user models.AuthUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, email, err := database.GetAuthData(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roles, err := database.GetUserRoles(id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error while scanning roles: " + err.Error()})
		return
	}
	// ! ПОЛУЧИТЬ РОЛИ И НАПИСАТЬ ФУНКЦИЮ ГЕНЕРАЦИИ ТОКЕНА

	token, err := jwtservice.GenerateToken(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, gin.H{"message": "Access granted"})
}
