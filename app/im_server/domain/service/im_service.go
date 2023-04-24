package service

import (
	"github.com/gin-gonic/gin"
	"github.com/penggehero/funny-im/app/im_server/domain/ws"
	log "github.com/sirupsen/logrus"
)

// IMService IM服务 todo
type IMService struct {
}

// NewIMService 创建IM服务
func NewIMService() *IMService {
	return &IMService{}

}

// ChatHandler 聊天接口
func (s *IMService) ChatHandler(context *gin.Context) {
	conn, err := ws.Upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Errorf("upgrade websocket error: %v", err)
		return
	}
	userId := context.Param("user_id")
	log.Infof("user [%s] connected.", userId)
	cc := &ws.ClientConn{
		Id:        userId,
		Conn:      conn,
		SendQueue: make(chan []byte),
	}
	ws.Pool.Add(cc)

	go cc.SendLoop()
	go cc.RecvLoop()

}

// LoginHandler 登录接口
func (s *IMService) LoginHandler(context *gin.Context) {

}

// RegisterHandler 注册接口
func (s *IMService) RegisterHandler(context *gin.Context) {

}

// UserDetailHandler 用户详情接口
func (s *IMService) UserDetailHandler(context *gin.Context) {

}

// CreateGroupChatHandler 创建群聊接口
func (s *IMService) CreateGroupChatHandler(context *gin.Context) {

}

// JoinGroupChatHandler 加入群聊接口
func (s *IMService) JoinGroupChatHandler(context *gin.Context) {

}

// QuitGroupChatHandler 退出群聊接口
func (s *IMService) QuitGroupChatHandler(context *gin.Context) {

}

// GroupChatMembersHandler 群聊成员接口
func (s *IMService) GroupChatMembersHandler(context *gin.Context) {

}

// GroupChatListHandler 群聊列表接口
func (s *IMService) GroupChatListHandler(context *gin.Context) {

}
