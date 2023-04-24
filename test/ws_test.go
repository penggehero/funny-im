package test

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/penggehero/funny-im/app/im_server/domain/ws"
)

func TestChat(t *testing.T) {

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/chat/1"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatal("dial:", err)
	}
	defer c.Close()

	message := &ws.Message{
		SenderID:   "1",
		Cmd:        0,
		ReceiverID: "2",
		Content:    "hello",
	}
	marshal, err := json.Marshal(message)
	if err != nil {
		t.Fatal("marshal:", err)
	}
	err = c.WriteMessage(websocket.TextMessage, marshal)

}
