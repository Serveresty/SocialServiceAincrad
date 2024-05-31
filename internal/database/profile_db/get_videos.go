package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"SocialServiceAincrad/utils"
	"context"
)

func GetVideosListById(id int) ([]models.Video, error) {
	rows, err := database.DB.Query(context.Background(), `SELECT s.video_id, s.title, s.created_at, s.views, s.preview_name FROM users_videos u JOIN videos s ON u.videos_list @> ARRAY[s.video_id] WHERE u.user_id = $1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		err := rows.Scan(&video.Id, &video.Title, &video.CreatedAt, &video.Views, &video.PreviewName)
		if err != nil {
			return nil, err
		}

		filePath := "../../storages/video_storage/previews/" + video.PreviewName + ".mp4.jpg"
		previewBase64, err := utils.ConvertToBase64(filePath)
		if err != nil {
			return nil, err
		}
		video.Preview = previewBase64
		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
