package main

import (
	"context"
	"flag"

	"github.com/shengchaohua/red-packet-backend/base/conf"
	datadm "github.com/shengchaohua/red-packet-backend/data/dm"
	datapkg "github.com/shengchaohua/red-packet-backend/data/pkg"
	"github.com/shengchaohua/red-packet-backend/infra"
	adminserver "github.com/shengchaohua/red-packet-backend/server_gin/server/admin"
	"github.com/shengchaohua/red-packet-backend/service"
)

var (
	configFilePath = flag.String("conf", "./conf/conf.toml", "admin app config file")
)

func init() {
	flag.Parse()
}

func main() {
	conf.InitAppConfig(*configFilePath)

	ctx := context.Background()

	// infra
	infra.InitInfra(ctx)

	// data
	datadm.InitDM()
	datapkg.InitPkg()

	// service
	service.InitService()

	// server
	adminserver.NewServer().Run()
}
