package userwalletpkg

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "userwalletpkg"

var (
	ErrCreateUserWallet        = errorpkg.New(pkgName, 1)
	ErrDeductUserWalletBalance = errorpkg.New(pkgName, 2)
	ErrAddUserWalletBalance    = errorpkg.New(pkgName, 3)

	// business
	ErrWalletBalanceNotEnough = errorpkg.New(pkgName, 10)
)
