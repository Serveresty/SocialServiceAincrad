package video

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func VideoGET(c *gin.Context) {
}

func VideoCurrentUserGET(c *gin.Context) {
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
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if claims.Subject == id {
		videos, err := profiledb.GetVideosListById(idInt)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": videos})
		return
	} else {
		privacy, err := profiledb.GetPrivacySettings(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if privacy.Video == "all" {
			videos, err := profiledb.GetVideosListById(idInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": videos})
			return
		}

		if privacy.Video == "friends" {
			currentUserId, err := strconv.Atoi(claims.Subject)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrorClaims.Error()})
				return
			}
			ok, err := profiledb.IsFriendOneByOne(idInt, currentUserId)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
				return
			}
			if !ok {
				c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
				return
			}
			videos, err := profiledb.GetVideosListById(idInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": videos})
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
			return
		}
	}
}
