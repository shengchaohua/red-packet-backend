package redpacketservice

import (
	redpacketpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/red_packet"
)

type defaultService struct {
	redPacketManager redpacketpkg.Manager
}

func NewDefaultService(
	redPacketManager redpacketpkg.Manager,
) Service {
	return &defaultService{
		redPacketManager: redPacketManager,
	}
}
