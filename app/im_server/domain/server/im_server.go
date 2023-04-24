package server

import (
	"github.com/gin-gonic/gin"
	"github.com/penggehero/funny-im/app/im_server/domain/service"
	"github.com/penggehero/funny-im/pkg/conf"
	"github.com/spf13/viper"
)

func init() {
	conf.Init()
}

// ImServer IM服务
type ImServer struct {
	engine  *gin.Engine
	service *service.IMService
}

// NewImServer 创建IM服务
func NewImServer(engine *gin.Engine, service *service.IMService) *ImServer {
	engine.Use(gin.Recovery())
	server := ImServer{engine, service}
	server.initRouter()
	return &server
}

// Run 运行服务
func (i *ImServer) Run() error {
	return i.engine.Run(viper.GetString("server.port"))
}

// initRouter 初始化路由
func (i *ImServer) initRouter() {
	apiGroup := i.engine.Group("/api")
	apiGroup.Any("/chat/:user_id", i.service.ChatHandler)

	userGroup := apiGroup.Group("/user")
	userGroup.POST("/login", i.service.LoginHandler)
	userGroup.POST("/register", i.service.RegisterHandler)
	userGroup.GET("/detail", i.service.UserDetailHandler)

	// 群聊接口
	groupChatGroup := apiGroup.Group("/group_chat")
	// 群聊创建接口
	groupChatGroup.POST("/create", i.service.CreateGroupChatHandler)
	// 群聊加入接口
	groupChatGroup.POST("/join", i.service.JoinGroupChatHandler)
	// 群聊退出接口
	groupChatGroup.POST("/quit", i.service.QuitGroupChatHandler)
	// 群聊成员列表接口
	groupChatGroup.GET("/members", i.service.GroupChatMembersHandler)
	// 群聊列表接口
	groupChatGroup.GET("/list", i.service.GroupChatListHandler)

}
