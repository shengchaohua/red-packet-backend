package constants

type Errcode uint32

// Errcode
// Note: use snake case naming for better readability
const (
	Errcode_Ok Errcode = 0

	// common
	Errcode_WrongParam Errcode = 1000
	Errcode_Server     Errcode = 1001

	// business
	Errcode_WalletBalanceNotEnough = 1100
	Errcode_UserNotInGroup         = 1101
)
