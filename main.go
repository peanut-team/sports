package main

import (
	"flag"
	"sports/config"
	"sports/internal"
	"sports/pkg/common/base"
	"sports/pkg/common/middleware"
	"sports/pkg/common/server"
	"sports/pkg/ctr"
	"sports/pkg/logger"
	"sports/pkg/notifier"
)

func main() {
	// set config path
	configPath := flag.String("c", base.ConfigPath, "配置文件路径")
	flag.Parse()
	port := flag.Int("p", base.ServerPort, "服务监听端口")
	flag.Parse()
	cfg := config.Config{}
	if err := server.ParseConfig(&cfg, *configPath); err != nil {
		panic("配置解析失败: " + err.Error())
	}
	if *port != base.ServerPort {
		cfg.Server.Port = *port
	}
	config.InitConfig(cfg)
	logger.InitLogger(&logger.Config{
		Debug: config.GetConfig().Debug,
	})

	// init db
	if err := ctr.InitDatabase(&cfg.Mysql); err != nil {
		panic("mysql配置加载失败：" + err.Error())
	}

	notifier.InitNotifier()


	// init gin
	engine := server.NewGinEngine(
		&cfg.Server,
		server.AddMiddleware(middleware.Cors()),
		internal.RouteApi,
	)
	engine.Start()
}
