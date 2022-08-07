package redpacketservice

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "redpacketservice"

var (
	ErrInvalidParams = errorpkg.New(pkgName, 1)
)
