package userwalletpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "userwalletpkg"

var (
	ErrDeductUserWallet = errorgrouppkg.New(pkgName, 1)
	ErrAddUserWallet    = errorgrouppkg.New(pkgName, 2)

	// business
	ErrWalletNotActive        = errorgrouppkg.New(pkgName, 10)
	ErrWalletBalanceNotEnough = errorgrouppkg.New(pkgName, 11)
)
