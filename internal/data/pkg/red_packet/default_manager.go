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

func (manager *defaultManager) CreateP2PRedPacket(
	ctx context.Context,
	session *xorm.Session,
	redPacketName string,
	redPacketResultType enum.RedPacketResultType,
	quantity uint32,
	amount uint32,
	receiverUserId uint64,
) (*redpacketmodel.RedPacket, error) {
	redPacket := &redpacketmodel.RedPacket{
		RedPacketTab: &redpacketmodel.RedPacketTab{
			RedPacketName:       redPacketName,
			RedPacketCategory:   enum.RedPacketCategoryP2P,
			RedPacketResultType: redPacketResultType,
			Quantity:            quantity,
			RemainingQuantity:   quantity,
			Amount:              amount,
		},
		ExtraData: &redpacketmodel.RedPacketExtraData{
			ReceiverUserId: receiverUserId,
		},
	}

	if err := manager.redPacketDM.InsertWithSession(ctx, session, redPacket); err != nil {
		return nil, ErrCreateRedPacket.WrapWithMsg(err, "[CreateP2PRedPacket]create_new_red_packet_error")
	}

	return redPacket, nil
}

func (manager *defaultManager) CreateGroupRedPacket(
	ctx context.Context,
	session *xorm.Session,
	redPacketName string,
	redPacketResultType enum.RedPacketResultType,
	quantity uint32,
	amount uint32,
	groupId uint64,
) (*redpacketmodel.RedPacket, error) {
	redPacket := &redpacketmodel.RedPacket{
		RedPacketTab: &redpacketmodel.RedPacketTab{
			RedPacketName:       redPacketName,
			RedPacketCategory:   enum.RedPacketCategoryGroup,
			RedPacketResultType: redPacketResultType,
			Quantity:            quantity,
			RemainingQuantity:   quantity,
			Amount:              amount,
		},
		ExtraData: &redpacketmodel.RedPacketExtraData{
			GroupId: groupId,
		},
	}

	if err := manager.redPacketDM.InsertWithSession(ctx, session, redPacket); err != nil {
		return nil, ErrCreateRedPacket.WrapWithMsg(err, "[CreateGroupRedPacket]create_new_red_packet_error")
	}

	return redPacket, nil
}
