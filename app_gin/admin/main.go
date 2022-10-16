package main

import (
	"flag"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	datadm "github.com/shengchaohua/red-packet-backend/internal/data/dm"
	datapkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg"
	internalpkg "github.com/shengchaohua/red-packet-backend/internal/pkg"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"github.com/shengchaohua/red-packet-backend/internal/service"
	adminserver "github.com/shengchaohua/red-packet-backend/server_gin/server/admin"
)

var (
	configFilePath = flag.String("conf", "./conf/test.toml", "admin app config file")
)

func init() {
	flag.Parse()
}

func main() {
	config.InitAppConfig(*configFilePath)

	// pkg
	logger.InitLogger(config.GetGlobalAppConfig().AdminConfig)
	ctx := logger.NewCtxWithTraceId()
	internalpkg.InitPkg(ctx)

	// data
	datadm.InitDM()
	datapkg.InitPkg()

	// service
	service.InitService()

	// server
	adminserver.NewServer().Run()
}
