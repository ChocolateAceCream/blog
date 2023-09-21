package service

import (
	"fmt"
	"strconv"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/library"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/request"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type NotificationService struct{}

const UNREAD = 2

func (es *CommentService) GetNotificationList(query request.NotificationCursorListParam, currentUser dbTable.User) (notificationBaseInfo []response.NotificationBaseInfo, total int64, err error) {
	db := global.DB.Model(&dbTable.Notification{}).Where("recipient_id  = ? ", currentUser.ID)
	notificationList := []dbTable.Notification{}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	fmt.Println("----total-----", total)
	db = db.Limit(query.PageSize)
	queryStr := "notifications.id > ?"
	if query.Desc {
		db = db.Order("id desc")
		queryStr = "notifications.id < ?"
	}
	if query.CursorId > 0 {
		db = db.Where(queryStr, query.CursorId)
	}
	if query.UnreadOnly {
		db = db.Where("status = ? ", UNREAD)
	}

	err = db.Preload("Initiator").Find(&notificationList).Error
	notificationBaseInfo = utils.MapSlice(notificationList, response.NotificationBaseInfoFormatter)
	return
}

func (ns *NotificationService) DeleteNotification(recipientID uint, notificationID int) (err error) {
	var notification dbTable.Notification
	notification.ID = uint(notificationID)
	if err = global.DB.Where("recipient_id = ? ", recipientID).Delete(&notification).Error; err != nil {
		return err
	}
	return nil
}

func (ns *NotificationService) ReadNotification(recipientID uint, notificationID int) error {
	q := global.DB.Model(&dbTable.Notification{}).Where("id = ? AND recipient_id  = ? AND status = ?", notificationID, recipientID, UNREAD).UpdateColumn("status", 1)
	if q.Error != nil {
		return q.Error
	}
	if q.RowsAffected == 1 {
		return q.UpdateColumn("status", 1).Error
	}
	return nil
}

func (ns *NotificationService) AddNotification(payload dbTable.Notification) {
	if err := global.DB.Create(&payload).Error; err != nil {
		global.LOGGER.Error("fail to create new notification", zap.Error(err))
		return
	}
	//send out message to mqtt
	library.PublishMqttMsg(fmt.Sprintf("notification%d", payload.RecipientID), strconv.FormatUint(uint64(payload.RecipientID), 10))
}

func (ns *NotificationService) GetUnreadCount(RecipientID uint) (count int64, err error) {
	const UNREAD = 2
	err = global.DB.Model(&dbTable.Notification{}).Where("recipient_id  = ? AND status = ?", RecipientID, UNREAD).Count(&count).Error
	return
}

func (ns *NotificationService) WSHandler(c *gin.Context, id uint) error {
	// sample code for print out request all headers
	// headers := c.Request.Header
	// for key, values := range headers {
	// 	for _, value := range values {
	// 		fmt.Printf("%s: %s\n", key, value)
	// 	}
	// }
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

		err = global.WS[uint(id)].Write(c, websocket.MessageText, []byte("New Notification"))
		if err != nil {
			global.LOGGER.Error("fail to send message through WS", zap.Error(err))
			return
		}
	}

	library.SubscribeMqttMsg(fmt.Sprintf("notification%d", id), f)

	for {
		// if conn has been reset by frontend
		if global.WS[id] != conn {
			return conn.Close(websocket.StatusNormalClosure, "Closing connection")
		}

		// Read a message from the WebSocket connection
		messageType, message, err := conn.Read(c)
		if err != nil {
			// Check if the error is related to a closed connection
			// if websocket.CloseStatus(err) == websocket.StatusNormalClosure || websocket.CloseStatus(err) == websocket.StatusGoingAway {
			// Handle the lost connection here
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
