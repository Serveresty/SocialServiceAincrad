package models

type ProfileData struct {
	Id         int       `json:"id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Sex        string    `json:"sex"`
	MoreInfo   UserInfo  `json:"more_info"`
	Social     Social    `json:"social"`
	Friends    []Friends `json:"friends"`
	Followers  []Friends `json:"followers"`
}

type UserInfo struct {
	ShortInfo       string `json:"short_info"`
	FamilyState     string `json:"family_state"`
	BornCity        string `json:"born_city"`
	CurrentResident string `json:"current_resident"`
}

type Social struct {
	PersonalWeb string `json:"personal_web"`
	Instagram   string `json:"instagram"`
	Steam       string `json:"steam"`
}

type Friends struct {
	FriendId        int    `json:"friend_id"`
	FriendUsername  string `json:"friend_username"`
	FriendFirstName string `json:"friend_first_name"`
	FriendLastName  string `json:"friend_last_name"`
}
