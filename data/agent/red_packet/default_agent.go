package redpacketagent

import redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"

type DefaultAgent struct {
	RedPacketDM redpacketdm.DataManager
}

var defaultAgent *DefaultAgent
