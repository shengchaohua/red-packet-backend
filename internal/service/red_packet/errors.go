package redpacketservice

import (
	errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"
)

const pkgName = "redpacketservice"

var (
	ErrWrongParam = errorgrouppkg.New(pkgName, 1)
)
