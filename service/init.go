package service

import (
	redpacketservice "github.com/shengchaohua/red-packet-backend/service/service/red_packet"
	userservice "github.com/shengchaohua/red-packet-backend/service/service/user"
)

func InitService() {
	redpacketservice.InitService()
	userservice.InitService()
}
