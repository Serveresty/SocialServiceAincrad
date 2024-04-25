package messagesdb

import (
	"SocialServiceAincrad/internal/database"
	"SocialServiceAincrad/models"
	"context"
	"strconv"
)

func GetMessages(sender string, receiver string) ([]models.Message, error) {
	senderInt, err := strconv.Atoi(sender)
	if err != nil {
		return nil, err
	}

	receiverInt, err := strconv.Atoi(receiver)
	if err != nil {
		return nil, err
	}

	rows, err := database.DB.Query(context.Background(), `SELECT u.first_name AS sender_name, uu.first_name AS receiver_name, m.message, m.created_at FROM messages m JOIN users_data u ON m.sender_id = u.user_id JOIN users_data uu ON m.receiver_id = uu.user_id WHERE (m.sender_id = $1 AND m.receiver_id = $2) OR (m.sender_id = $3 AND m.receiver_id = $4);`, senderInt, receiverInt, receiverInt, senderInt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.Sender, &msg.Receiver, &msg.Messages, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}
