package signin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Sign-In Page"})
}
