package redpacketagent

import (
	"context"

	redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"
	"github.com/shengchaohua/red-packet-backend/data/enum"
	redpacketmodel "github.com/shengchaohua/red-packet-backend/data/model/red_packet"
	"xorm.io/xorm"
)

type defaultAgent struct {
	RedPacketDM redpacketdm.DataManager
}

func NewDefaultAgent(redPacketDM redpacketdm.DataManager) Agent {
	return &defaultAgent{
		RedPacketDM: redPacketDM,
	}
}

func (agent *defaultAgent) CreateRedPacket(
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
			RedPacketName: redPacketName,
		},

		RedPacketCategory: redPacketCategory,
		RedPacketType:     redPacketType,
	}

	if err := agent.RedPacketDM.Insert(ctx, session, redPacket); err != nil {
		return nil, ErrCreateRedPacket.WrapWithMsg(err, "[CreateRedPacket]failed to create red packet")
	}

	return redPacket, nil
}
