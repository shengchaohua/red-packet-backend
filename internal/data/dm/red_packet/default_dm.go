package redpacketdm

import (
	"context"
	"fmt"

	"xorm.io/xorm"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
)

type defaultDM struct {
	database.EngineManager
}

func NewDefaultDM(engineManager database.EngineManager) DM {
	return &defaultDM{
		EngineManager: engineManager,
	}
}

func (dm *defaultDM) Insert(
	ctx context.Context,
	session *xorm.Session,
	redPacket *redpacketmodel.RedPacket,
) error {
	if session == nil {
		session = dm.GetMasterEngine().Table(redpacketmodel.RedPacketTableName)
	}

	affected, err := session.Table(redpacketmodel.RedPacketTableName).InsertOne(redPacket)
	if err != nil {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf(
			"insert_db_error|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
			redPacket.RedPacketName, redPacket.RedPacketCategory, redPacket.RedPacketType,
		))
	}
	if affected == 0 {
		return ErrInsert.WithMsg(fmt.Sprintf(
			"insert_db_failed|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
			redPacket.RedPacketName, redPacket.RedPacketCategory, redPacket.RedPacketType,
		))
	}

	return nil
}

func (dm *defaultDM) QueryById(
	ctx context.Context,
	redPacketId uint64,
	querySlave bool,
	queryMaster bool,
) (*redpacketmodel.RedPacket, error) {
	if !querySlave && !queryMaster {
		return nil, ErrParam.WithMsg("both querySlave and queryMaster are false")
	}

	var (
		redPacket = &redpacketmodel.RedPacket{}
		err       error
	)

	if querySlave {
		redPacket, err = dm.loadById(ctx, redPacketId, false)
		if err != nil {
			return nil, err
		}
		if redPacket != nil {
			return redPacket, nil
		}
	}

	if queryMaster {
		redPacket, err = dm.loadById(ctx, redPacketId, true)
		if err != nil {
			return nil, err
		}
	}

	return redPacket, nil
}

func (dm *defaultDM) loadById(
	ctx context.Context,
	redPacketId uint64,
	queryMaster bool,
) (*redpacketmodel.RedPacket, error) {
	var (
		redPacket = &redpacketmodel.RedPacket{}
		engine    = dm.GetSlaveEngine()
	)
	if queryMaster {
		engine = dm.GetMasterEngine()
	}

	has, err := engine.Table(redpacketmodel.RedPacketTableName).
		Where("red_packet_id = ?", redPacketId).
		Get(redPacket)
	if err != nil {
		return nil, ErrQuery.WrapWithMsg(err, fmt.Sprintf(
			"query_db_error|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
			redPacket.RedPacketName, redPacket.RedPacketCategory, redPacket.RedPacketType,
		))
	}
	if !has {
		return nil, nil
	}

	return redPacket, nil
}
