package userservice

import (
	"github.com/shengchaohua/red-packet-backend/internal/constants"
	errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"
)

const pkgName = "userservice"

var (
	ErrWrongParam = errorpkg.New(pkgName, int(constants.Errcode_WrongParam))
	ErrServer     = errorpkg.New(pkgName, int(constants.Errcode_Server))
)
