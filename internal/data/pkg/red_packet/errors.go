package redpacketpkg

import errpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "redpacketpkg"

var (
	ErrCreateRedPacket = errpkg.New(pkgName, 1)
)
