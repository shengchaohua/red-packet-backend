package usergrouprelationpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "usergrouprelationpkg"

var (
	ErrCheckUserInGroup = errorgrouppkg.New(pkgName, 1)
)
