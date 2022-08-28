package redpacketpkg

import (
	"context"

	"xorm.io/xorm"

	redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"
	"github.com/shengchaohua/red-packet-backend/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
)

type Manager interface {
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

var defaultManagerInstance Manager

func InitManager() {
	defaultDM := redpacketdm.GetDefaultDM()
	defaultManagerInstance = NewDefaultAgent(defaultDM)
}

func GetDefaultManager() Manager {
	return defaultManagerInstance
}
