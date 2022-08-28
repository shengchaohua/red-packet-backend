package redpacketdm

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "redpacketdm"

// errros
var (
	ErrInsert = errorgrouppkg.New(pkgName, 1)
	ErrQuery  = errorgrouppkg.New(pkgName, 2)
)
