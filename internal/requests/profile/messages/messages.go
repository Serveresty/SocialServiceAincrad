package messages

import (
	jwtservice "SocialServiceAincrad/internal/jwt-service"
	"SocialServiceAincrad/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешаем любые origin
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// type Clients struct {
// 	mx      sync.Mutex
// 	clients map[string]*websocket.Conn
// }

// var clients = &Clients{clients: make(map[string]*websocket.Conn)}
var clients = make(map[string]*websocket.Conn)

func ChatGET(c *gin.Context) {
	// fmt.Println("TY312312312312")
	// err := utils.CheckAlreadyToken(c)
	// if err == nil {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Println("TYT1")
	// claims, err := jwtservice.ParseToken(c)
	// if err != nil {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	// 	return
	// }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	var msg map[string]string
	if err := conn.ReadJSON(&msg); err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	claims, err := jwtservice.ParseTokenString(msg["authToken"])
	if err != nil {
		//c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	id := msg["id"]

	if id == "" {
		//Get all chats
		//c.JSON(http.StatusOK, gin.H{})
		return
	}
	// _, err = messagesdb.GetMessages(claims.Subject, id)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	//c.JSON(http.StatusFound, gin.H{"message": msgs})

	clients[claims.Subject] = conn
	// clients.mx.Lock()
	// clients.clients[claims.Subject] = conn
	// clients.mx.Unlock()
	fmt.Println("TYT1")
	fmt.Println(clients)
	for {
		// Пример чтения данных от клиента
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("TYT2")
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
		if client, ok := clients[id]; ok {
			if err := client.WriteMessage(messageType, p); err != nil {
				fmt.Println("TYT3")
			}
		}
		if client, ok := clients[claims.Subject]; ok {
			if err := client.WriteMessage(messageType, p); err != nil {
				fmt.Println("TYT3")
			}
		}
	}
}
