package agent

import (
	redpacketagent "github.com/shengchaohua/red-packet-backend/data/agent/red_packet"
	useragent "github.com/shengchaohua/red-packet-backend/data/agent/user"
)

func InitAgent() {
	redpacketagent.InitAgent()
	useragent.InitAgent()
}
