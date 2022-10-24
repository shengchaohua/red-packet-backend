package redpacketpkg

import (
	"context"

	"xorm.io/xorm"

	redpacketdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet"
	"github.com/shengchaohua/red-packet-backend/internal/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/internal/data/model/red_packet"
)

type defaultManager struct {
	redPacketDM redpacketdm.DM
}

func NewDefaultAgent(redPacketDM redpacketdm.DM) Manager {
	return &defaultManager{
		redPacketDM: redPacketDM,
	}
}

func (manager *defaultManager) CreateRedPacket(
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

	if err := manager.redPacketDM.InsertWithSession(ctx, session, redPacket); err != nil {
		return nil, ErrCreateRedPacket.WrapWithMsg(err, "[CreateRedPacket]create_new_red_packet_error")
	}

	return redPacket, nil
}
