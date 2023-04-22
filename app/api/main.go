package main

import (
	"context"
	"flag"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpacketdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet"
	userdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user"
	usergrouprelationgdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_group_relation"
	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
	userwallettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet_transaction"
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	userpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user"
	usergrouprelationgpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_group_relation"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	userwallettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
	userservice "github.com/shengchaohua/red-packet-backend/internal/service/user"
	"github.com/shengchaohua/red-packet-backend/server/server"
)

var (
	configFilePath = flag.String("conf", "./conf/api/dev.toml", "api config file for local dev env")
	port           = flag.String("port", "8020", "server port")
)

func main() {
	flag.Parse()
	config.InitConfig(*configFilePath)

	// pkg
	ctx := context.Background()
	logger.InitLogger(config.GetGlobalConfig().ServerConfig)
	ctx = logger.NewCtxWithTraceId(ctx, "main")
	database.InitEngineManager(ctx)

	// data dm
	redpacketdm.InitDM()
	userdm.InitDM()
	usergrouprelationgdm.InitDM()
	userwalletdm.InitDM()
	userwallettxndm.InitDM()

	// data pkg
	redpacketpkg.InitManager()
	userpkg.InitManager()
	usergrouprelationgpkg.InitManager()
	userwalletpkg.InitManager()
	userwallettxnpkg.InitManager()

	// service
	redpacketservice.InitService()
	userservice.InitService()

	// server
	server.NewApiServer().Run(*port)
}
