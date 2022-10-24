package main

import (
	"flag"

	"github.com/shengchaohua/red-packet-backend/internal/config"
	redpacketdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet"
	userwalletdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet"
	userwallettxndm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user_wallet_transaction"
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	userwallettxnpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet_transaction"
	internalpkg "github.com/shengchaohua/red-packet-backend/internal/pkg"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
	apiserver "github.com/shengchaohua/red-packet-backend/server_gin/server/api"
)

var (
	configFilePath = flag.String("conf", "./conf/test.toml", "api app config file")
)

func init() {
	flag.Parse()
}

func main() {
	config.InitAppConfig(*configFilePath)

	// pkg
	logger.InitLogger(config.GetGlobalAppConfig().APIConfig)
	ctx := logger.NewCtxWithTraceId()
	internalpkg.InitPkg(ctx)

	// data dm
	redpacketdm.InitRedPacketDM()
	userwalletdm.InitUserWalletDM()
	userwallettxndm.InitUserWalletTxnDM()

	// data pkg
	redpacketpkg.InitRedPacketManager()
	userwalletpkg.InitUserWalletManager()
	userwallettxnpkg.InitUserWalletTxnManager()

	// service
	redpacketservice.InitRedPacketService()

	// server
	apiserver.NewServer().Run()
}
