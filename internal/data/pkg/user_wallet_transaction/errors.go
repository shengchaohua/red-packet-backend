package redpackettxnpkg

import errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"

const pkgName = "userwallettxnpkg"

var (
	ErrAddUserWalletTxn = errorpkg.New(pkgName, 1)
)
