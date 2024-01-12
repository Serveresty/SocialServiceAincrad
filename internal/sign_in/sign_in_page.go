package signin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func SignInGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Sign-In Page"})
}

// POST
func SignInPOST(c *gin.Context) {

}
