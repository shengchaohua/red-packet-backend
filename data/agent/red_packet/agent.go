package redpacketagent

import redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"

type Agent interface {
}

func InitAgent() {
	redPakcetDM := redpacketdm.GetDataManager()
	defaultAgent = &DefaultAgent{
		RedPacketDM: redPakcetDM,
	}
}

func GetAgent() Agent {
	return defaultAgent
}
