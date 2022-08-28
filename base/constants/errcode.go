package constants

type Errcode uint32

// Errcode
// use snake case naming for better readability
const (
	Errcode_Ok     Errcode = 0
	Errcode_Param  Errcode = 1
	Errcode_Server Errcode = 2
)
