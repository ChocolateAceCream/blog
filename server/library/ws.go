package library

import "nhooyr.io/websocket"

type ActiveWS map[string]*websocket.Conn

func InitWS() ActiveWS {
	return make(ActiveWS)
}
