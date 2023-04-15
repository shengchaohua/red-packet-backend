package redpacketpkg

import (
	"context"

	"xorm.io/xorm"

	redpacketdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
)

type Manager interface {
	CreateP2PRedPacket(
		ctx context.Context,
		session *xorm.Session,
		redPacketName string,
		redPacketResultType enum.RedPacketResultType,
		quantity uint32,
		amount uint32,
		receiverUserId uint64,
	) (*redpacketmodel.RedPacket, error)

	CreateGroupRedPacket(
		ctx context.Context,
		session *xorm.Session,
		redPacketName string,
		redPacketResultType enum.RedPacketResultType,
		quantity uint32,
		amount uint32,
		groupId uint64,
	) (*redpacketmodel.RedPacket, error)
}

var defaultManagerInstance Manager

func InitManager() {
	defaultDM := redpacketdm.GetDM()
	defaultManagerInstance = NewDefaultAgent(defaultDM)
}

func GetManager() Manager {
	if defaultManagerInstance == nil {
		panic("defaultManagerInstance has not been inited")
	}
	return defaultManagerInstance
}
