package messages

import (
	messagesdb "SocialServiceAincrad/internal/database/messages_db"
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[string]*websocket.Conn)

func ChatGET(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	var msg map[string]string
	if err := conn.ReadJSON(&msg); err != nil {
		return
	}

	claims, err := jwtservice.ParseTokenString(msg["authToken"])
	if err != nil {
		return
	}

	id := msg["id"]

	if id == "" {
		//Get all chats
		//c.JSON(http.StatusOK, gin.H{})
		return
	}
	msgg, err := messagesdb.GetMessages(claims.Subject, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := conn.WriteJSON(msgg); err != nil {
		return
	}

	clients[claims.Subject] = conn

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("error while read message")
			return
		}

		var data map[string]string
		err = json.Unmarshal([]byte(p), &data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		message, ok := data["message"]
		if !ok {
			fmt.Println("Message key not found in JSON")
			return
		}

		msg := models.Message{Sender: claims.Subject, Receiver: id, Messages: message, CreatedAt: time.Now()}
		p, err = json.Marshal(&msg)
		if err != nil {
			continue
		}
		err = messagesdb.CreateMessage(msg)
		if err != nil {
			log.Println("error save in database message " + err.Error())
		}
		if client, ok := clients[id]; ok {
			if err := client.WriteMessage(messageType, p); err != nil {
				log.Println("receiver not online")
			}
		}
		if client, ok := clients[claims.Subject]; ok {
			if err := client.WriteMessage(messageType, p); err != nil {
				log.Println("sender not online")
			}
		}
	}
}
