package usergrouprelationpkg

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "usergrouprelationpkg"

var (
	ErrCheckUserInGroup = errorpkg.New(pkgName, 1)
)
