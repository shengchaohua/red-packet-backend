package redpacketdm

import (
	"context"

	"xorm.io/xorm"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/infra/database"
)

type DataManager interface {
	Insert(
		ctx context.Context,
		session *xorm.Session,
		redPacket *redpacketmodel.RedPacket,
	) error

	LoadById(
		ctx context.Context,
		session *xorm.Session,
		redPacketId uint64,
	) (*redpacketmodel.RedPacket, error)
}

var defaultDMInstance DataManager

func InitDM() {
	defaultDBEngine := database.GetDefaultDBEngineManager()
	defaultDMInstance = NewDefaultDM(defaultDBEngine)
}

func GetDefaultDM() DataManager {
	return defaultDMInstance
}
