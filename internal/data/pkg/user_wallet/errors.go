package userwalletpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "userwalletpkg"

var (
	ErrDeductUserWallet = errorgrouppkg.New(pkgName, 1)
)
