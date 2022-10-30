package userwalletpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "userwalletpkg"

var (
	ErrCreateUserWallet        = errorgrouppkg.New(pkgName, 1)
	ErrDeductUserWalletBalance = errorgrouppkg.New(pkgName, 2)
	ErrAddUserWalletBalance    = errorgrouppkg.New(pkgName, 3)

	// business
	ErrWalletNotActive        = errorgrouppkg.New(pkgName, 10)
	ErrWalletBalanceNotEnough = errorgrouppkg.New(pkgName, 11)
)
