package redpacketdm

import (
	"context"
	"fmt"

	"xorm.io/xorm"

	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
	"github.com/shengchaohua/red-packet-backend/infra/database"
)

type defaultDM struct {
	database.EngineManager
}

func NewDefaultDM(engineManager database.EngineManager) DataManager {
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
		session = dm.GetMasterEngine().Table(redpacketmodel.TableRedPacket)
	}

	affected, err := session.Table(redpacketmodel.TableRedPacket).InsertOne(redPacket)
	if err != nil {
		return ErrInsert.WrapWithMsg(err, fmt.Sprintf(
			"[Insert]fail to insert new red packet|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
			redPacket.RedPacketName, redPacket.RedPacketCategory, redPacket.RedPacketType,
		))
	}
	if affected == 0 {
		return ErrInsert.WithMsg(fmt.Sprintf(
			"[Insert]insert nothing|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
			redPacket.RedPacketName, redPacket.RedPacketCategory, redPacket.RedPacketType,
		))
	}

	return nil
}

func (dm *defaultDM) LoadById(
	ctx context.Context,
	redPacketId uint64,
	querySlave bool,
	queryMaster bool,
) (*redpacketmodel.RedPacket, error) {
	var redPacket = &redpacketmodel.RedPacket{}

	if querySlave {
		has, err := dm.GetSlaveEngine().Table(redpacketmodel.TableRedPacket).Where("red_packet_id = ?", redPacketId).Get(redPacket)
		if err != nil {
			return nil, ErrQuery.WrapWithMsg(err, fmt.Sprintf(
				"[LoadById]fail to get red packet from slave db|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
				redPacket.RedPacketName, redPacket.RedPacketCategory, redPacket.RedPacketType,
			))
		}
		if has {
			return redPacket, nil
		}
	}

	if queryMaster {
		has, err := dm.GetMasterEngine().Table(redpacketmodel.TableRedPacket).Where("red_packet_id = ?", redPacketId).Get(redPacket)
		if err != nil {
			return nil, ErrQuery.WrapWithMsg(err, fmt.Sprintf(
				"[LoadById]fail to get red packet from master db|red_packet_name=%s,red_packet_category=%s,red_packet_type=%s",
				redPacket.RedPacketName, redPacket.RedPacketCategory, redPacket.RedPacketType,
			))
		}
		if has {
			return redPacket, nil
		}
	}

	return redPacket, nil
}
