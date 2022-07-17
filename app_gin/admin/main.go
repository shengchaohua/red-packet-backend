package main

import (
	"context"
	"flag"

	"github.com/shengchaohua/red-packet-backend/common/conf"
	"github.com/shengchaohua/red-packet-backend/data/agent"
	"github.com/shengchaohua/red-packet-backend/data/dm"
	"github.com/shengchaohua/red-packet-backend/infra"
	adminroute "github.com/shengchaohua/red-packet-backend/server_gin/route/admin"
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
	dm.InitDM()
	agent.InitAgent()

	// service
	service.InitService()

	// server
	adminroute.NewRouter().Run()
}
