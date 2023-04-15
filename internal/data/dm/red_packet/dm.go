package redpacketdm

import (
	"context"
	"xorm.io/xorm"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type DM interface {
	InsertWithSession(
		ctx context.Context,
		session *xorm.Session,
		redPacket *redpacketmodel.RedPacket,
	) error

	LoadByIdWithSession(
		ctx context.Context,
		session *xorm.Session,
		redPacketId uint64,
	) (*redpacketmodel.RedPacket, error)

	LoadById(
		ctx context.Context,
		redPacketId uint64,
		querySlave bool,
		queryMaster bool,
	) (*redpacketmodel.RedPacket, error)
}

var (
	defaultDMInstance DM
)

func InitDM() {
	mainDBEngineManager := database.GetMainDBEngineManager()
	defaultDMInstance = NewDefaultDM(mainDBEngineManager)
}

func GetDM() DM {
	if defaultDMInstance == nil {
		panic("defaultDMInstance is nil")
	}
	return defaultDMInstance
}
