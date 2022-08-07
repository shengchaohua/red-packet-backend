package redpacketagent

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "redpacketagent"

var (
	ErrCreateRedPacket = errorpkg.New(pkgName, 1)
)
