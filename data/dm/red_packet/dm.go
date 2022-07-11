package redpacketdm

import (
	"context"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
)

type DataManager interface {
	Create(ctx context.Context, redPacket *redpacketmodel.RedPacket) error
}

var (
	defaultDM DataManager
)

func InitDataManger() {
	defaultDM = &DefaultDM{}
}

func GetDataManager() DataManager {
	return defaultDM
}
