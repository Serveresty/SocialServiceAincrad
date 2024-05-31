package profiledb

import (
	"SocialServiceAincrad/internal/database"
	"context"
)

func SetPreviewToVideo(filename string) error {
	_, err := database.DB.Exec(context.Background(), `
	UPDATE videos SET preview_name = $1 WHERE filename = $2;
	`, filename, filename)
	if err != nil {
		return err
	}
	return nil
}
