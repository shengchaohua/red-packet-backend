package redpacketpkg

import (
	"context"

	"xorm.io/xorm"

	redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"
	"github.com/shengchaohua/red-packet-backend/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
)

type defaultManager struct {
	RedPacketDM redpacketdm.DataManager
}

func NewDefaultAgent(redPacketDM redpacketdm.DataManager) Manager {
	return &defaultManager{
		RedPacketDM: redPacketDM,
	}
}

func (agent *defaultManager) CreateRedPacket(
	ctx context.Context,
	session *xorm.Session,
	redPacketName string,
	redPacketCategory enum.RedPacketCategory,
	redPacketType enum.RedPacketType,
	quantity uint32,
	amount uint32,
) (*redpacketmodel.RedPacket, error) {
	redPacket := &redpacketmodel.RedPacket{
		RedPacketTab: &redpacketmodel.RedPacketTab{
			RedPacketName:     redPacketName,
			RedPacketCategory: redPacketCategory,
			RedPacketType:     redPacketType,
			Quantity:          quantity,
			RemainingQuantity: quantity,
			Amount:            amount,
		},
	}

	if err := agent.RedPacketDM.Insert(ctx, session, redPacket); err != nil {
		return nil, ErrCreateRedPacket.WrapWithMsg(err, "[CreateRedPacket]failed to create new red packet")
	}

	return redPacket, nil
}
