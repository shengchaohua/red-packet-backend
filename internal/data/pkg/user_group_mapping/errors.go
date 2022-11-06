package usergroupmappingpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "usergroupmappingpkg"

var (
	ErrCheckUserInGroup = errorgrouppkg.New(pkgName, 1)
)
