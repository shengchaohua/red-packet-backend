package redpacketdm

import (
	"context"
	"fmt"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/database"
	"github.com/shengchaohua/red-packet-backend/internal/pkg/logger"
	"xorm.io/xorm"
)

type defaultDM struct {
	database.EngineManager
	tableName string
}

func NewDefaultDM(engineManager database.EngineManager) DM {
	return &defaultDM{
		EngineManager: engineManager,
		tableName:     redpacketmodel.RedPacketTableName,
	}
}

func (dm *defaultDM) InsertWithSession(
	ctx context.Context,
	session *xorm.Session,
	redPacket *redpacketmodel.RedPacket,
) error {
	if session == nil {
		return ErrParam.WithMsg("session cannot be nil")
	}

	return dm.insert(ctx, session, redPacket)
}

func (dm *defaultDM) insert(
	ctx context.Context,
	session *xorm.Session,
	redPacket *redpacketmodel.RedPacket,
) error {
	if session == nil {
		logger.Logger(ctx).Info("[defaultDM.insert]session_is_nil,use_default_session")
		session = dm.GetMasterEngine().Table(dm.tableName)
	}

	redPacketTab, err := redPacket.ModelToTab()
	if err != nil {
		return ErrData.Wrap(err)
	}

	affected, err := session.Table(dm.tableName).InsertOne(redPacketTab)
	if err != nil {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf(
			"insert_db_error|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
			redPacket.RedPacketName,
			redPacket.RedPacketCategory.String(),
			redPacket.RedPacketResultType.String(),
		))
	}
	if affected == 0 {
		return ErrInsert.WithMsg(fmt.Sprintf(
			"insert_db_failed|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
			redPacket.RedPacketName,
			redPacket.RedPacketCategory.String(),
			redPacket.RedPacketResultType.String(),
		))
	}

	return nil
}

func (dm *defaultDM) LoadByIdWithSession(
	ctx context.Context,
	session *xorm.Session,
	redPacketId uint64,
) (*redpacketmodel.RedPacket, error) {
	if session == nil {
		return nil, ErrParam.WithMsg("session cannot be nil")
	}

	return dm.loadById(ctx, session, redPacketId, true)
}

func (dm *defaultDM) LoadById(
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
		redPacket, err = dm.loadById(ctx, nil, redPacketId, false)
		if err != nil {
			return nil, err
		}
		if redPacket != nil {
			return redPacket, nil
		}
	}

	if queryMaster {
		redPacket, err = dm.loadById(ctx, nil, redPacketId, true)
		if err != nil {
			return nil, err
		}
		if redPacket == nil {
			return nil, nil
		}
	}

	return redPacket, nil
}

func (dm *defaultDM) loadById(
	ctx context.Context,
	session *xorm.Session,
	redPacketId uint64,
	queryMaster bool,
) (*redpacketmodel.RedPacket, error) {
	var (
		redPacketTab = &redpacketmodel.RedPacketTab{}
		engine       = dm.GetSlaveEngine()
	)
	if queryMaster {
		engine = dm.GetMasterEngine()
	}
	if session == nil {
		logger.Logger(ctx).Info("[defaultDM.loadById]session_is_nil,use_new_session")
		session = engine.Table(dm.tableName)
	}

	has, err := engine.Table(dm.tableName).
		Where("red_packet_id = ?", redPacketId).
		Get(redPacketTab)
	if err != nil {
		return nil, ErrQuery.WrapWithMsg(err, fmt.Sprintf("query db error|red_packet_id=%d", redPacketId))
	}
	if !has {
		return nil, nil
	}

	return redPacketTab.TabToModel()
}
