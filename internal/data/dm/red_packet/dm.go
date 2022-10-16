package redpacketdm

import (
	"context"

	"xorm.io/xorm"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type DM interface {
	Insert(
		ctx context.Context,
		session *xorm.Session,
		redPacket *redpacketmodel.RedPacket,
	) error

	QueryById(
		ctx context.Context,
		redPacketId uint64,
		querySlave bool,
		queryMaster bool,
	) (*redpacketmodel.RedPacket, error)
}

var (
	defaultDMInstance DM
)

func InitRedPacketDM() {
	defaultDBEngine := database.GetMainDBEngineManager()
	defaultDMInstance = NewDefaultDM(defaultDBEngine)
}

func GetRedPacketDM() DM {
	return defaultDMInstance
}
