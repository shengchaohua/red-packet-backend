package userpkg

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "userpkg"

var (
	ErrCreateUser = errorpkg.New(pkgName, 1)
)
