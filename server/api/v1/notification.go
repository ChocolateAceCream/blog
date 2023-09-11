package apiV1

import (
	"fmt"
	"strconv"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type NotificationApi struct{}

func (a *NotificationApi) WSHandler(c *gin.Context) {

	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		global.LOGGER.Error("User not logged in", zap.Error(err))
		response.FailWithUnauthorized("User not logged in", c)
		return
	}
	fmt.Println("----WSHandler-------")

	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // will not verify the server's certificate and will trust any certificate presented by the server.
	})

	global.WS[strconv.FormatUint(uint64(currentUser.ID), 10)] = conn

	if err != nil {
		global.LOGGER.Error("fail to create WS", zap.Error(err))
		return
	}
	err = utils.SetValueToSession(c, "websocket", conn)
	fmt.Println("----SetValueToSession err-------", err)

	defer conn.Close(websocket.StatusInternalError, "the sky is falling")
	fmt.Println("-----connn origin-----", &conn)
	f := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("msg.Payload: %s\n", msg.Payload())
		// Send a message to the client
		err = global.WS[string(msg.Payload())].Write(c, websocket.MessageText, []byte("New Message"))
		if err != nil {
			global.LOGGER.Error("fail to send message through WS", zap.Error(err))
			return
		}
	}

	library.SubscribeMqttMsg(strconv.FormatUint(uint64(currentUser.ID), 10), f)

	for {
		// Read a message from the WebSocket connection
		_, _, err := conn.Read(c)
		fmt.Println("---------c--------", c)
		if err != nil {
			// Check if the error is related to a closed connection
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure || websocket.CloseStatus(err) == websocket.StatusGoingAway {
				// Handle the lost connection here
				global.LOGGER.Error("WebSocket connection lost", zap.Error(err))
				library.UnsubscribeMqttMsg(strconv.FormatUint(uint64(currentUser.ID), 10))
				delete(global.WS, strconv.FormatUint(uint64(currentUser.ID), 10))
				fmt.Println(len(global.WS))
				break
			}
		}
	}
}
