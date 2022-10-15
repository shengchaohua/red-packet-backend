package main

import (
	"context"
	"flag"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	datadm "github.com/shengchaohua/red-packet-backend/internal/data/dm"
	datapkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg"
	internalpkg "github.com/shengchaohua/red-packet-backend/internal/pkg"
	"github.com/shengchaohua/red-packet-backend/internal/service"
	apiserver "github.com/shengchaohua/red-packet-backend/server_gin/server/api"
)

var (
	configFilePath = flag.String("conf", "./conf/conf.toml", "admin app config file")
)

func init() {
	flag.Parse()
}

func main() {
	config.InitAppConfig(*configFilePath)

	ctx := context.Background()

	// infra
	internalpkg.InitPkg(ctx)

	// data
	datadm.InitDM()
	datapkg.InitPkg()

	// service
	service.InitService()

	// server
	apiserver.NewServer().Run()
}
