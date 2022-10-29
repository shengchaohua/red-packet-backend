package redpackethandler

import (
	"github.com/shengchaohua/red-packet-backend/internal/constants"
	"github.com/shengchaohua/red-packet-backend/server_gin/route"
)

var (
	allowedErrorMap = map[string]map[constants.Errcode]bool{
		route.RouteCreateRedPacket: {
			constants.Errcode_WalletNotAvtive:        true,
			constants.Errcode_WalletBalanceNotEnough: true,
		},

		route.RouteOpenRedPacket: {},
	}
)
