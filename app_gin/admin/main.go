package main

import (
	"flag"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpacketdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet"
	redpackettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet_transaction"
	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	redpackettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet_transaction"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
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
	database.InitEngineManager(ctx)

	// data dm
	redpacketdm.InitRedPacketDM()
	userwalletdm.InitUserWalletDM()
	redpackettxndm.InitRedPacketTxnDM()

	// data pkg
	redpacketpkg.InitRedPacketManager()
	userwalletpkg.InitUserWalletManager()
	redpackettxnpkg.InitRedPacketTxnManager()

	// service
	redpacketservice.InitRedPacketService()

	// server
	adminserver.NewServer().Run()
}
