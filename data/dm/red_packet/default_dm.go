package redpacketdm

import (
	"context"

	"xorm.io/xorm"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/infra/database"
)

type defaultDM struct {
	database.EngineManager
}

func NewDefaultDM(engineManager database.EngineManager) *defaultDM {
	return &defaultDM{
		EngineManager: engineManager,
	}
}

func (dm *defaultDM) Create(
	ctx context.Context,
	session *xorm.Session,
	redPacket *redpacketmodel.RedPacket,
) error {
	if session == nil {
		session = dm.GetMasterEngine().Table(redpacketmodel.TableRedPacket)
	}

	affected, err := session.Table(redpacketmodel.TableRedPacket).InsertOne(redPacket)
	if err != nil {
		return err
	}
	if affected == 0 {
		return err
	}

	return nil
}
