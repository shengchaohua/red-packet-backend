package redpackethandler

import (
	"github.com/shengchaohua/red-packet-backend/internal/constants"
	"github.com/shengchaohua/red-packet-backend/server_gin/route"
)

var (
	allowedErrorMap = map[string]map[constants.Errcode]bool{
		route.RouteCreateRedPacket: {
			constants.Errcode_WalletBalanceNotEnough: true,
			constants.Errcode_UserNotInGroup:         true,
		},

		route.RouteOpenRedPacket: {},
	}
)
