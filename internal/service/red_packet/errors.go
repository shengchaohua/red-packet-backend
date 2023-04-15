package redpacketservice

import (
	"github.com/shengchaohua/red-packet-backend/internal/constants"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	errorpkg "github.com/shengchaohua/red-packet-backend/pkg/error_pkg"
)

const pkgName = "redpacketservice"

var (
	ErrWrongParam = errorpkg.New(pkgName, int(constants.Errcode_WrongParam))
	ErrServer     = errorpkg.New(pkgName, int(constants.Errcode_Server))

	ErrWalletBalanceNotEnough = errorpkg.New(pkgName, int(constants.Errcode_WalletBalanceNotEnough))
	ErrUserNotInGroup         = errorpkg.New(pkgName, int(constants.Errcode_UserNotInGroup))
)

const (
	errmsgWalletBalanceNotEnough = "user wallet balance is not enough"
)

func errorMapping(err error) error {
	if userwalletpkg.ErrWalletBalanceNotEnough.Is(err) {
		return ErrWalletBalanceNotEnough.WrapWithMsg(err, errmsgWalletBalanceNotEnough)
	}
	return err
}
