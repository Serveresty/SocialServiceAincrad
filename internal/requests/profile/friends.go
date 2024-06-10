package profile

import (
	cerr "SocialServiceAincrad/custom_errors"
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FriendsGET(c *gin.Context) {
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

	currentUserId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrorClaims.Error()})
		return
	}

	userID := c.Query("id")
	section := c.Query("section")
	if userID == "" {
		friendData, err := showBySection(c, section, currentUserId)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": friendData})
		return
	} else {
		idInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		privacy, err := profiledb.GetPrivacySettings(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if privacy.Friends == "all" {
			friendData, err := showBySection(c, section, idInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": friendData})
			return
		}
		if privacy.Friends == "friends" {
			ok, err := profiledb.IsFriendOneByOne(idInt, currentUserId)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
				return
			}
			if !ok {
				c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
				return
			}
			friendData, err := showBySection(c, section, idInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": friendData})
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": cerr.ErrNoAccessByPrivacy.Error()})
			return
		}
	}
}

func showBySection(c *gin.Context, section string, id int) ([]models.Friends, error) {
	if section == "all" || section == "friends" {
		friends, err := profiledb.GetFriends(id)
		if err != nil {
			return nil, err
		}
		return friends, nil
	}
	if section == "followers" {
		followers, err := profiledb.GetFollowers(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, err
		}
		return followers, nil
	}
	if section == "" {
		friends, err := profiledb.GetFriends(id)
		if err != nil {
			return nil, err
		}
		return friends, nil
	}
	return nil, cerr.ErrBadRequest
}
