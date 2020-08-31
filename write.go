package ws_common

import (
	"github.com/gorilla/websocket"
)

func WriteMessage(c *websocket.Conn, messageType []byte, messageContent interface{}) error {
	dst, err := encode(messageType, messageContent)
	if err != nil {
		return err
	}
	return c.WriteMessage(websocket.BinaryMessage, dst)
}
