package dm

import (
	redpacketdm "github.com/shengchaohua/red-packet-backend/data/dm/red_packet"
	userdm "github.com/shengchaohua/red-packet-backend/data/dm/user"
)

func InitDM() {
	redpacketdm.InitDM()
	userdm.InitDM()
}
