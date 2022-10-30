package redpacketservice

import (
	"github.com/shengchaohua/red-packet-backend/internal/constants"
	userwalletpkg "github.com/shengchaohua/red-packet-backend/internal/data/pkg/user_wallet"
	errorgrouppkg "github.com/shengchaohua/red-packet-backend/pkg/error_group"
)

const pkgName = "redpacketservice"

var (
	ErrWrongParam             = errorgrouppkg.New(pkgName, int(constants.Errcode_WrongParam))
	ErrServer                 = errorgrouppkg.New(pkgName, int(constants.Errcode_Server))
	ErrWalletNotActive        = errorgrouppkg.New(pkgName, int(constants.Errcode_WalletNotAvtive))
	ErrWalletBalanceNotEnough = errorgrouppkg.New(pkgName, int(constants.Errcode_WalletBalanceNotEnough))
)

const (
	errmsgWalletBalanceNotEnough = "user wallet balance is not enough"
)

func errorMapping(err error) error {
	if userwalletpkg.ErrWalletBalanceNotEnough.Is(err) {
		return ErrWalletBalanceNotEnough.WrapWithMsg(err, errmsgWalletBalanceNotEnough)
	}
	return ErrServer.Wrap(err)
}
