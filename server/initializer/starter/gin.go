package starter

import (
	"fmt"
	"red-server/controller"
	"red-server/global"
	"red-server/middleware"

	"github.com/gin-gonic/gin"
)

type GinStarter struct {
	BaseStarter
	engine *gin.Engine
}

func (s *GinStarter) Name() string {
	return "api 服务"
}

func (s *GinStarter) Init() {
	s.engine = gin.Default()
}

func (s *GinStarter) Setup() {
	// 注册路由
	api := s.engine.Group("/api/v1")
	controller.NewAuthController().RegisterRoute(api)
	api.Use(middleware.Auth())
	controller.NewUserController().RegisterRoute(api)
	controller.NewAccountController().RegisterRoute(api)
	controller.NewEnvelopeGoodsController().RegisterRoute(api)
	controller.NewEnvelopeGoodsItemController().RegisterRoute(api)
}

func (s *GinStarter) Start() {
	global.Logger.Info("启动 gin...")
	config := global.CONFIG.System

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	global.Logger.Infof("Start server on: %s", addr)
	s.engine.Run(addr)
}

func (s *GinStarter) StartBlocking() bool {
	return true
}
