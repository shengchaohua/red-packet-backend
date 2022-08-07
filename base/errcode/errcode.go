package errcode

type Errcode uint32

// Errcode
// use snake case naming for better readability
const (
	Errcode_ok     Errcode = 0
	Errcode_server Errcode = 1
)
