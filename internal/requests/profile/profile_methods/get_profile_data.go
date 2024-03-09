package profilemethods

import (
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	"SocialServiceAincrad/models"
	"strconv"
)

func GetProfileData(id string, privacy models.PrivacySettings) (*models.ProfileData, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	data, err := profiledb.GetProfileById(intId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
