package constants

type Errcode uint32

// Errcode
// Note: use snake case naming for better readability
const (
	Errcode_Ok         Errcode = 0
	Errcode_WrongParam Errcode = 1000
	Errcode_Server     Errcode = 1001
)

var errCodeMap = map[Errcode]bool{
	Errcode_WrongParam: true,
	Errcode_Server:     true,
}

func ParseErrcodeEnum(code int) Errcode {
	errcode := Errcode(code)
	if _, ok := errCodeMap[errcode]; ok {
		return errcode
	}
	return Errcode_Server
}
