package service

import (
	redpacketservice "github.com/shengchaohua/red-packet-backend/service/red_packet"
	userservice "github.com/shengchaohua/red-packet-backend/service/user"
)

func InitService() {
	redpacketservice.InitService()
	userservice.InitService()
}
