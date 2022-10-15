package datadm

import (
	redpacketdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/red_packet"
	userdm "github.com/shengchaohua/red-packet-backend/internal/data/dm/user"
)

func InitDM() {
	redpacketdm.InitRedPacketDM()
	userdm.InitUserDM()
}
