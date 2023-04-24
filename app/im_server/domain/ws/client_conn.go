package ws

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	SingleMsg = iota
	RoomMsg
)

type ClientConn struct {
	ctx       context.Context
	Id        string
	Conn      *websocket.Conn
	SendQueue chan []byte
	quit      chan struct{}
}

func (c *ClientConn) Dispatch(message []byte) {
	var msg Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Errorf("unmarshal error: %v", err)
	}
	switch msg.Cmd {
	case SingleMsg:
		c.Send(msg.ReceiverID, message)
	case RoomMsg:
		// TODO

	}
}

func (c *ClientConn) Send(userId string, message []byte) {
	cilent, ok := Pool.clients.Load(userId)
	if !ok {
		log.Errorf("user %s not online", userId)
		return
	}
	cilent.(*ClientConn).SendQueue <- message
}

func (c *ClientConn) SendLoop() {
	for {
		send := <-c.SendQueue
		err := c.Conn.WriteMessage(websocket.TextMessage, send)
		if err != nil {
			log.Errorf("write error: %v", err)
			Pool.Remove(c)
			return
		}
	}
}

func (c *ClientConn) RecvLoop() {
	for {
		messageType, data, err := c.Conn.ReadMessage()
		if err != nil {
			log.Errorf("read error: %v", err)
			Pool.Remove(c)
			return
		}
		switch messageType {
		case websocket.TextMessage:
			c.Dispatch(data)
		}
	}

}
