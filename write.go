package ws_common

import (
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	recws "github.com/maurodelazeri/gorilla-reconnect"
)

//带自动重连功能发消息
func WriteMessageWithAutoConnect(c *recws.RecConn, messageType []byte, messageContent interface{}) (err error) {
	defer func() {
		if exception := recover(); exception != nil {
			err = errors.New(fmt.Sprintf("WriteMessage异常:%v", exception))
			return
		}
	}()
	if !c.IsConnected() {
		time.Sleep(100 * time.Millisecond)
		err = errors.New("WriteMessage错误:连接已断开")
		return
	}
	dst, err := encode(messageType, messageContent)
	if err != nil {
		return err
	}
	return c.WriteMessage(websocket.BinaryMessage, dst)
}

//发消息
func WriteMessage(c *websocket.Conn, messageType []byte, messageContent interface{}) error {
	dst, err := encode(messageType, messageContent)
	if err != nil {
		return err
	}
	return c.WriteMessage(websocket.BinaryMessage, dst)
}
