package ws

// Message 消息结构体
type Message struct {
	Id         int64  `json:"id,omitempty" form:"id"`                   // 消息ID
	SenderID   string `json:"sender_id,omitempty" form:"sender_id"`     // 发送者ID
	Cmd        int    `json:"cmd,omitempty" form:"cmd"`                 // 消息类型
	ReceiverID string `json:"receiver_id,omitempty" form:"receiver_id"` // 接收者ID
	Content    string `json:"content,omitempty" form:"content"`         // 消息内容
}
