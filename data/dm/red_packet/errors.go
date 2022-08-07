package redpacketdm

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "redpacketdm"

// errros
var (
	ErrInsert = errorpkg.New(pkgName, 1)
)
