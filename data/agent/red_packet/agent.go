package redpacketagent

import redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"

type Agent interface {
}

func InitAgent() {
	defaultDM := redpacketdm.GetDefaultDM()
	defaultAgent = &DefaultAgent{
		RedPacketDM: defaultDM,
	}
}

func GetAgent() Agent {
	return defaultAgent
}
