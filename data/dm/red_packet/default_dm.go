package redpacketdm

import (
	"context"

	"github.com/go-xorm/xorm"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
)

type DefaultDM struct {
}

func (dm *DefaultDM) Create(
	ctx context.Context,
	session *xorm.Session,
	redPacket *redpacketmodel.RedPacket,
) error {
	if session == nil {
		session = dm.
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
