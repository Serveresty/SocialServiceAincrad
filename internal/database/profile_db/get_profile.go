package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func GetProfileById(id int) (*models.ProfileData, error) {
	var profileData models.ProfileData
	profileData.Id = id

	err := GetGeneralInfo(&profileData)
	if err != nil {
		return nil, err
	}

	err = GetAdditionalInfo(&profileData)
	if err != nil {
		return nil, err
	}

	friends, err := GetFriends(id)
	if err != nil {
		return nil, err
	}

	followers, err := GetFollowers(id)
	if err != nil {
		return nil, err
	}

	profileData.Friends = friends
	profileData.Followers = followers

	return &profileData, nil
}

func GetGeneralInfo(profileData *models.ProfileData) error {
	row1 := database.DB.QueryRow(context.Background(), `SELECT first_name, last_name, sex FROM users_data WHERE user_id = $1`, profileData.Id)
	err := row1.Scan(&profileData.First_name, &profileData.Last_name, &profileData.Sex)
	if err != nil {
		return err
	}
	return nil
}

func GetAdditionalInfo(profileData *models.ProfileData) error {
	row2 := database.DB.QueryRow(context.Background(), `SELECT i.short_info, i.family_state, i.born_city, i.current_resident, s.personal_web, s.instagram, s.steam FROM info i JOIN social s ON i.user_id = $1 and s.user_id = i.user_id`, profileData.Id)
	err := row2.Scan(&profileData.MoreInfo.ShortInfo, &profileData.MoreInfo.FamilyState, &profileData.MoreInfo.BornCity, &profileData.MoreInfo.CurrentResident, &profileData.Social.PersonalWeb, &profileData.Social.Instagram, &profileData.Social.Steam)
	if err != nil {
		return err
	}
	return nil
}

func GetFriends(id int) ([]models.Friends, error) {
	var friends []models.Friends
	row3, err := database.DB.Query(context.Background(), `SELECT ud.user_id, ud.first_name, ud.last_name, ud.username FROM users_data ud JOIN friends f ON CASE WHEN f.first = $1 THEN ud.user_id = f.second WHEN f.second = $1 THEN ud.user_id = f.first END JOIN friend_status fs ON fs.status_name = $2 AND fs.status_id = f.status_id`, id, "friend")
	if err != nil {
		return nil, err
	}

	for row3.Next() {
		var friend models.Friends
		err = row3.Scan(&friend.FriendId, &friend.FriendFirstName, &friend.FriendLastName, &friend.FriendUsername)
		if err != nil {
			return nil, err
		}

		friends = append(friends, friend)
	}
	return friends, nil
}

func GetFollowers(id int) ([]models.Friends, error) {
	var follows []models.Friends
	row4, err := database.DB.Query(context.Background(), `SELECT ud.user_id, ud.first_name, ud.last_name, ud.username FROM users_data ud JOIN friends f ON f.second = $1 JOIN friend_status fs ON fs.status_name = $2 AND fs.status_id = f.status_id`, id, "follower")
	if err != nil {
		return nil, err
	}

	for row4.Next() {
		var friend models.Friends
		err = row4.Scan(&friend.FriendId, &friend.FriendFirstName, &friend.FriendLastName, &friend.FriendUsername)
		if err != nil {
			return nil, err
		}

		follows = append(follows, friend)
	}
	return follows, nil
}

func IsFriendOneByOne(firstId int, secondId int) (bool, error) {
	var status string
	row4, err := database.DB.Query(context.Background(), `SELECT status_name FROM friends JOIN friend_status ON friends.status_id = friend_status.status_id WHERE first = $1 AND second = $2 UNION SELECT status_name FROM friends JOIN friend_status ON friends.status_id = friend_status.status_id WHERE first = $2 AND second = $1;`, firstId, secondId)
	if err != nil {
		return false, err
	}
	err = row4.Scan(&status)
	if err != nil {
		return false, err
	}
	return true, nil
}
