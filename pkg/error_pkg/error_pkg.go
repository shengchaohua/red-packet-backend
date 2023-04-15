package errorpkg

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorPkg defines an error type with package, code and message
type ErrorPkg struct {
	cause error
	pkg   string
	code  int
	msg   string // pure error msg
}

// New returns a new ErrorPkg whose error is nil.
func New(pkg string, code int) *ErrorPkg {
	return &ErrorPkg{
		pkg:  pkg,
		code: code,
	}
}

// Error returns error string
func (errorPkg *ErrorPkg) Error() string {
	if errorPkg.cause == nil {
		if errorPkg.msg == "" {
			return fmt.Sprintf("ErrorPkg[Pkg='%s',Code=%d]", errorPkg.pkg, errorPkg.code)
		}

		return fmt.Sprintf("ErrorPkg[Pkg='%s',Code=%d,Msg='%s']",
			errorPkg.pkg,
			errorPkg.code,
			errorPkg.msg,
		)
	}
	return errorPkg.cause.Error()
}

// WithMsg appends a message to an ErrorPkg whose cause is nil.
func (errorPkg *ErrorPkg) WithMsg(msg string) *ErrorPkg {
	if errorPkg.cause != nil {
		return errorPkg
	}

	return &ErrorPkg{
		pkg:  errorPkg.pkg,
		code: errorPkg.code,
		msg:  msg,
	}
}

// Wrap wraps an error.
// Note: use WrapWithMsg as much as possible
func (errorPkg *ErrorPkg) Wrap(err error) *ErrorPkg {
	errmsg := errorPkg.msg
	if errorPkg.msg == "" {
		errmsg = err.Error()
	}

	return &ErrorPkg{
		cause: errors.Wrap(err, errorPkg.Error()),
		pkg:   errorPkg.pkg,
		code:  errorPkg.code,
		msg:   errmsg,
	}
}

// WrapWithMsg wraps an error with a message.
func (errorPkg *ErrorPkg) WrapWithMsg(err error, msg string) *ErrorPkg {
	return errorPkg.WithMsg(msg).Wrap(err)
}

// getErrorPkgCode gets pkg and code from an error.
func as(err error) (*ErrorPkg, bool) {
	errGroup, ok := err.(*ErrorPkg)
	return errGroup, ok
}

// Is checks if a given err is a specific ErrorPkg
func (errorPkg *ErrorPkg) Is(err error) bool {
	if asErrGroup, ok := as(err); ok {
		return asErrGroup.pkg == errorPkg.pkg && asErrGroup.code == errorPkg.code
	}
	return false
}

// GetErrcode gets error code from an error.
func GetErrcode(err error) (int, bool) {
	errGroup, ok := as(err)
	if ok {
		return errGroup.code, ok
	}
	return 0, ok
}

// GetErrmsg gets pure error message from an error.
func GetErrmsg(err error) string {
	errGroup, ok := as(err)
	if ok {
		return errGroup.msg
	}
	return err.Error()
}
