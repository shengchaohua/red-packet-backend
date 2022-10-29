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

func errorMapping(err error) error {
	if userwalletpkg.ErrWalletNotActive.Is(err) {
		return ErrWalletNotActive.Wrap(err)
	}
	if userwalletpkg.ErrWalletBalanceNotEnough.Is(err) {
		return ErrWalletBalanceNotEnough.Wrap(err)
	}
	return ErrServer.Wrap(err)
}
