package signup

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Sign-Up Page"})
}
