package audio

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AudioGET(c *gin.Context) {
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

	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if claims.Subject == id {
		songs, err := profiledb.GetAudiosListById(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": songs})
		return
	} else {
		privacy, err := profiledb.GetPrivacySettings(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if privacy.Audio == "all" {
			songs, err := profiledb.GetAudiosListById(idInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": songs})
			return
		}

		if privacy.Audio == "friends" {
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
			songs, err := profiledb.GetAudiosListById(idInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": songs})
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
			return
		}
	}
}
