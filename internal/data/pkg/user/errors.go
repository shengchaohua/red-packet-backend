package userpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "userpkg"

var (
	ErrCreateUser = errorgrouppkg.New(pkgName, 1)
)
