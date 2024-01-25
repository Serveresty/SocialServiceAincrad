package profile

import (
	profiledb "SocialServiceAincrad/internal/database/profile_db"
	"SocialServiceAincrad/models"
	"strconv"
)

func GetProfileData(id string) (*models.ProfileData, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		data, err := profiledb.GetProfileByUsername(id)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	data, err := profiledb.GetProfileById(intId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
