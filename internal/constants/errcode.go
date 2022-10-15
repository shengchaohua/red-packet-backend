package constants

type Errcode uint32

// Errcode
// Note: use snake case naming for better readability
const (
	Errcode_Ok Errcode = 0

	Errcode_WrongParam Errcode = 1
	Errcode_Server     Errcode = 2
)
