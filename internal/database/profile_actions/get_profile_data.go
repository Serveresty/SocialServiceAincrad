package profileactions

import (
	"SocialServiceAincrad/models"
	"strconv"
)

func GetProfileData(id string) (*models.ProfileData, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		data, err := GetProfileByUsername(id)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	data, err := GetProfileById(intId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
