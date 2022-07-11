package main

import (
	"flag"

	"github.com/shengchaohua/red-packet-backend/common/config"
	"github.com/shengchaohua/red-packet-backend/data/agent"
	"github.com/shengchaohua/red-packet-backend/data/dm"
	"github.com/shengchaohua/red-packet-backend/pkg"
	adminroute "github.com/shengchaohua/red-packet-backend/server_gin/route/admin"
	"github.com/shengchaohua/red-packet-backend/service"
)

var (
	configFilePath = flag.String("conf", "", "admin app config file")
)

func init() {
	flag.Parse()
}

func main() {
	router := adminroute.InitRouter()

	config.InitAppConfig(*configFilePath)

	// pkg
	pkg.InitPkg()

	// data
	dm.InitDataManager()
	agent.InitAgent()

	// service
	service.InitService()

	if err := router.Run(); err != nil {
		panic("fail to run admin app")
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
