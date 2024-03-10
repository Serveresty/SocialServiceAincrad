package profile

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FriendsGET(c *gin.Context) {
	err := utils.CheckAlreadyToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}

	claims, err := jwtservice.ParseToken(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	currentUserId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrorClaims.Error()})
		return
	}

	userID := c.Query("id")
	idInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	section := c.Query("section")
	if userID == "" {
		showBySection(c, section, currentUserId)
		return
	} else {
		privacy, err := profiledb.GetPrivacySettings(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if privacy.Friends == "all" {
			showBySection(c, section, idInt)
			return
		} else if privacy.Friends == "friends" {
			ok, err := profiledb.IsFriendOneByOne(idInt, currentUserId)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
				return
			}
			if !ok {
				c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
				return
			}
			showBySection(c, section, idInt)
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
			return
		}
	}
}

func showBySection(c *gin.Context, section string, id int) {
	if section == "all" || section == "friends" {
		friends, err := profiledb.GetFriends(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": friends})
		return
	}
	if section == "followers" {
		followers, err := profiledb.GetFollowers(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": followers})
		return
	}
}
