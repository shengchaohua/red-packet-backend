package main

import (
	"flag"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpacketdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet"
	userdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user"
	usergroupmappingdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_group_mapping"
	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
	userwallettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet_transaction"
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	userpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user"
	usergroupmappingpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_group_mapping"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	userwallettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
	userservice "github.com/shengchaohua/red-packet-backend/internal/service/user"
	server "github.com/shengchaohua/red-packet-backend/server_gin/server"
)

var (
	configFilePath = flag.String("conf", "./conf/admin/test.toml", "admin config file")
)

func main() {
	flag.Parse()
	config.InitAppConfig(*configFilePath)

	// pkg
	logger.InitLogger(config.GetGlobalAppConfig().ServerConfig)
	ctx := logger.NewCtxWithTraceId()
	database.InitEngineManager(ctx)

	// data dm
	redpacketdm.InitDM()
	userdm.InitDM()
	usergroupmappingdm.InitDM()
	userwalletdm.InitDM()
	userwallettxndm.InitDM()

	// data pkg
	redpacketpkg.InitManager()
	userpkg.InitManager()
	usergroupmappingpkg.InitManager()
	userwalletpkg.InitManager()
	userwallettxnpkg.InitManager()

	// service
	redpacketservice.InitService()
	userservice.InitService()

	// server
	server.NewAdminServer().Run()
}
