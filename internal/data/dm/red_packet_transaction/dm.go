package userwallettxndm

import (
	"context"

	"xorm.io/xorm"

	redpackettxnmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet_transaction"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type DM interface {
	InsertWithSession(
		ctx context.Context,
		session *xorm.Session,
		redPacketTransaction *redpackettxnmodel.RedPacketTransaction,
	) error
}

var (
	defaultDMInstance DM
)

func InitRedPacketTxnDM() {
	defaultDBEngine := database.GetMainDBEngineManager()
	defaultDMInstance = NewDefaultDM(defaultDBEngine)
}

func GetRedPacketTxnDM() DM {
	return defaultDMInstance
}
