package userwalletdm

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "userwalletdm"

// errros
var (
	ErrParam  = errorpkg.New(pkgName, 1)
	ErrData   = errorpkg.New(pkgName, 2)
	ErrInsert = errorpkg.New(pkgName, 3)
	ErrQuery  = errorpkg.New(pkgName, 4)
	ErrUpdate = errorpkg.New(pkgName, 5)
)
