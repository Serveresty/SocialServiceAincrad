package utils

import (
	cerr "SocialServiceAincrad/custom_errors"

	"github.com/gin-gonic/gin"
)

func CheckAlreadyToken(c *gin.Context) error {
	token := c.GetHeader("Authorization")

	if token != "" {
		return cerr.AlreadyAuthorized
	}

	return nil
}
