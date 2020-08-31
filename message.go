package ws_common

import (
	"bytes"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/ha666/ws-common/protocol"
)

var (
	MESSAGEPING         = []byte{112, 105, 110, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MESSAGEPONG         = []byte{112, 111, 110, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MESSAGEREAD         = []byte{114, 101, 97, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MESSAGEWRITE        = []byte{119, 114, 105, 116, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MESSAGESUBSCRIPTION = []byte{115, 117, 98, 115, 99, 114, 105, 112, 116, 105, 111, 110, 0, 0, 0, 0}
	MESSAGEPUBLISH      = []byte{112, 117, 98, 108, 105, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

func encode(messageType []byte, messageContent interface{}) ([]byte, error) {
	var (
		dst []byte
		err error
	)
	if bytes.Compare(messageType, MESSAGEPING) == 0 {
		value, ok := messageContent.(*protocol.Ping)
		if !ok || value == nil {
			return nil, errors.New("解析Ping请求出错")
		}
		dst, err = proto.Marshal(value)
	} else if bytes.Compare(messageType, MESSAGEPONG) == 0 {
		value, ok := messageContent.(*protocol.Pong)
		if !ok || value == nil {
			return nil, errors.New("解析Pong请求出错")
		}
		dst, err = proto.Marshal(value)
	} else if bytes.Compare(messageType, MESSAGEREAD) == 0 {
		value, ok := messageContent.(*protocol.Read)
		if !ok || value == nil {
			return nil, errors.New("解析Read请求出错")
		}
		dst, err = proto.Marshal(value)
	} else if bytes.Compare(messageType, MESSAGEWRITE) == 0 {
		value, ok := messageContent.(*protocol.Write)
		if !ok || value == nil {
			return nil, errors.New("解析Write请求出错")
		}
		dst, err = proto.Marshal(value)
	} else if bytes.Compare(messageType, MESSAGESUBSCRIPTION) == 0 {
		value, ok := messageContent.(*protocol.Subscription)
		if !ok || value == nil {
			return nil, errors.New("解析Subscription请求出错")
		}
		dst, err = proto.Marshal(value)
	} else if bytes.Compare(messageType, MESSAGEPUBLISH) == 0 {
		value, ok := messageContent.(*protocol.Publish)
		if !ok || value == nil {
			return nil, errors.New("解析Publish请求出错")
		}
		dst, err = proto.Marshal(value)
	} else {
		return nil, errors.New("无效的消息类型")
	}
	if err != nil {
		return nil, errors.New("序列化请求出错:" + err.Error())
	}
	return bytesCombine(messageType, dst), nil
}

func decode(message []byte) (dst proto.Message, messageType []byte, err error) {
	messageType = message[0:16]
	messageVal := message[16:]
	if bytes.Compare(messageType, MESSAGEPING) == 0 {
		dst = &protocol.Ping{}
	} else if bytes.Compare(messageType, MESSAGEPONG) == 0 {
		dst = &protocol.Pong{}
	} else if bytes.Compare(messageType, MESSAGEREAD) == 0 {
		dst = &protocol.Read{}
	} else if bytes.Compare(messageType, MESSAGEWRITE) == 0 {
		dst = &protocol.Write{}
	} else if bytes.Compare(messageType, MESSAGESUBSCRIPTION) == 0 {
		dst = &protocol.Subscription{}
	} else if bytes.Compare(messageType, MESSAGEPUBLISH) == 0 {
		dst = &protocol.Publish{}
	} else {
		return dst, nil, errors.New("无效的消息类型")
	}
	err = proto.Unmarshal(messageVal, dst)
	if err != nil {
		return nil, nil, err
	}
	return dst, messageType, nil
}

func bytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
