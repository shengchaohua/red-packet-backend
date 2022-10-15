package service

import (
	redpacketservice "github.com/shengchaohua/red-packet-backend/internal/service/red_packet"
	userservice "github.com/shengchaohua/red-packet-backend/internal/service/user"
)

func InitService() {
	redpacketservice.InitRedPacketService()
	userservice.InitService()
}
