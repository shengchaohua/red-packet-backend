package redpacketpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "redpacketpkg"

var (
	ErrCreateRedPacket = errorgrouppkg.New(pkgName, 1)
)
