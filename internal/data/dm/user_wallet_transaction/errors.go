package userwallettxndm

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "userwalletdm"

// errros
var (
	ErrParam  = errorgrouppkg.New(pkgName, 1)
	ErrData   = errorgrouppkg.New(pkgName, 2)
	ErrInsert = errorgrouppkg.New(pkgName, 3)
	ErrQuery  = errorgrouppkg.New(pkgName, 4)
	ErrUpdate = errorgrouppkg.New(pkgName, 5)
)
