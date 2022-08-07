package redpacketagent

import (
	"context"

	"xorm.io/xorm"

	redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"
	"github.com/shengchaohua/red-packet-backend/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
)

type Agent interface {
	CreateRedPacket(
		ctx context.Context,
		session *xorm.Session,
		redPacketName string,
		redPacketCategory enum.RedPacketCategory,
		redPacketType enum.RedPacketType,
		quantity uint32,
		amount uint32,
	) (*redpacketmodel.RedPacket, error)
}

var defaultAgentInstance Agent

func InitAgent() {
	defaultDM := redpacketdm.GetDefaultDM()
	defaultAgentInstance = NewDefaultAgent(defaultDM)
}

func GetDefaultAgent() Agent {
	return defaultAgentInstance
}
