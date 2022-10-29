package userservice

import (
	"github.com/shengchaohua/red-packet-backend/internal/constants"
	errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"
)

const pkgName = "userservice"

var (
	ErrWrongParam = errorgrouppkg.New(pkgName, int(constants.Errcode_WrongParam))
	ErrServer     = errorgrouppkg.New(pkgName, int(constants.Errcode_Server))
)