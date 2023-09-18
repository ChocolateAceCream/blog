package service

import (
	"fmt"
	"strconv"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type NotificationService struct{}

func (ns *NotificationService) WSHandler(c *gin.Context, id uint) error {
	headers := c.Request.Header

	for key, values := range headers {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // will not verify the server's certificate and will trust any certificate presented by the server.
	})

	global.WS[id] = conn

	if err != nil {
		return err
	}

	defer conn.Close(websocket.StatusInternalError, "the sky is falling")
	f := func(client mqtt.Client, msg mqtt.Message) {
		// Send a message to the client
		str := string(msg.Payload())
		id, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			global.LOGGER.Error("fail to parse mqtt msg to socket it", zap.Error(err))
			return
		}

		err = global.WS[uint(id)].Write(c, websocket.MessageText, []byte("New Message"))
		if err != nil {
			global.LOGGER.Error("fail to send message through WS", zap.Error(err))
			return
		}
	}

	library.SubscribeMqttMsg(fmt.Sprintf("notification%d", id), f)

	for {
		// Read a message from the WebSocket connection
		messageType, message, err := conn.Read(c)
		if err != nil {
			fmt.Println("-------1--err from socket---", err)
			// Check if the error is related to a closed connection
			// if websocket.CloseStatus(err) == websocket.StatusNormalClosure || websocket.CloseStatus(err) == websocket.StatusGoingAway {
			// Handle the lost connection here
			fmt.Println("-------2--err from socket---", err)
			library.UnsubscribeMqttMsg(fmt.Sprintf("notification%d", id))
			delete(global.WS, id)
			fmt.Println(len(global.WS))
			// }
			return err
		}
		// handle ping message from client and send back pong response
		if messageType == websocket.MessageText && string(message) == "ping" {
			if err := conn.Write(c, websocket.MessageText, []byte("pong")); err != nil {
				// Handle error
				return err
			}
		}
	}
}
