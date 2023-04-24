//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/penggehero/funny-im/app/im_server/domain/server"
	"github.com/penggehero/funny-im/app/im_server/domain/service"
)

func newIMServer() *server.ImServer {
	wire.Build(
		gin.New,
		service.NewIMService,
		server.NewImServer,
	)
	return &server.ImServer{}
}
