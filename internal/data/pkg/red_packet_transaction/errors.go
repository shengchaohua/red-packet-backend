package redpackettxnpkg

import errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"

const pkgName = "userwallettxnpkg"

var (
	ErrAddUserWalletTxn = errorgrouppkg.New(pkgName, 1)
)
