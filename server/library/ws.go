package library

import "nhooyr.io/websocket"

type ActiveWS map[uint]*websocket.Conn

func InitWS() ActiveWS {
	return make(ActiveWS)
}
