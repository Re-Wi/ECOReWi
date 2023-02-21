package initialize

import (
	ecoRouter "BackendGoECOReWi/apps/economy/router"
	logRouter "BackendGoECOReWi/apps/log/router"
	"BackendGoECOReWi/pkg/global"
	"BackendGoECOReWi/pkg/middleware"
	"BackendGoECOReWi/pkg/transport"
)

func InitRouter() *transport.HttpServer {
	// server配置
	serverConfig := global.Conf.Server
	server := transport.NewHttpServer(serverConfig.GetPort())

	container := server.Container
	// 是否允许跨域
	if serverConfig.Cors {
		container.Filter(middleware.Cors(container).Filter)
	}
	// 流量限制
	if serverConfig.Rate.Enable {
		container.Filter(middleware.Rate)
	}
	//日志系统
	{
		logRouter.InitOperLogRouter(container)
	}
	//经济管理
	{
		ecoRouter.InitDynamicEcoRouter(container)
	}
	// api接口
	middleware.SwaggerConfig(container)
	return server
}
