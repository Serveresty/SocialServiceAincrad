package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
)

func GetPrivacySettings(id int) (*models.PrivacySettings, error) {
	var privacy models.PrivacySettings
	row := database.DB.QueryRow(context.Background(),
		`SELECT 
			sp.status_name as saved_photos, 
			g.status_name as groups, 
			a.status_name as audio, 
			v.status_name as video, 
			f.status_name as friends, 
			p.status_name as posts, 
			c.status_name as comments, 
			m.status_name as messages 
		FROM privacy_settings ps 
		JOIN privacy_statuses sp ON sp.status_id = ps.saved_photos 
		JOIN privacy_statuses g ON g.status_id = ps.groups
		JOIN privacy_statuses a ON a.status_id = ps.audio
		JOIN privacy_statuses v ON v.status_id = ps.video
		JOIN privacy_statuses f ON f.status_id = ps.friends
		JOIN privacy_statuses p ON p.status_id = ps.posts
		JOIN privacy_statuses c ON c.status_id = ps.comments
		JOIN privacy_statuses m ON m.status_id = ps.messages
		WHERE ps.user_id = $1;`, 1)

	err := row.Scan(&privacy.SavedPhotos, &privacy.Groups, &privacy.Audio, &privacy.Video, &privacy.Friends, &privacy.Posts, &privacy.Comments, &privacy.Messages)
	if err != nil {
		return &models.PrivacySettings{}, err
	}

	return &privacy, nil
}
