package ws_common

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/maurodelazeri/gorilla-reconnect"
)

func ReadMessage(c *recws.RecConn) (dst proto.Message, messageType []byte, err error) {
	defer func() {
		if exception := recover(); exception != nil {
			log.Printf("出现异常:%v", exception)
			err = errors.New(fmt.Sprintf("ReadMessage异常:%v", exception))
			return
		}
	}()
	if !c.IsConnected() {
		log.Println("连接已断开")
		time.Sleep(100 * time.Millisecond)
		err = errors.New("ReadMessage错误:连接已断开")
		return
	}
	_, message, err := c.ReadMessage()
	if err != nil {
		time.Sleep(30 * time.Millisecond)
		err = errors.New("ReadMessage错误:" + err.Error())
		return
	}
	if message == nil || len(message) < 16 {
		return nil, nil, errors.New("无效的消息")
	}
	return decode(message)
}
