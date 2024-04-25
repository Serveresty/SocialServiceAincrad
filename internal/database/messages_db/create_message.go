package messagesdb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
	"strconv"
)

func CreateMessage(data models.Message) error {
	senderInt, err := strconv.Atoi(data.Sender)
	if err != nil {
		return err
	}

	receiverInt, err := strconv.Atoi(data.Receiver)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec(context.Background(), `INSERT INTO "messages" (sender_id, receiver_id, message, created_at) VALUES($1,$2,$3,$4)`,
		senderInt, receiverInt, data.Messages, data.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
