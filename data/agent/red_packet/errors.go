package redpacketagent

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "redpacketagent"

var (
	ErrCreateRedPacket = errorgrouppkg.New(pkgName, 1)
)
