package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"context"
)

func AddViewToVideo(vid string) error {
	_, err := database.DB.Exec(context.Background(), `
	UPDATE videos SET views = views + 1 WHERE video_id = $1;
	`, vid)
	if err != nil {
		return err
	}
	return nil
}
