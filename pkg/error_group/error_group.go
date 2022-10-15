package errorgrouppkg

import (
	"fmt"

	"github.com/pkg/errors"
)

// errorGroup defines an error type with package, code and message
type errorGroup struct {
	cause error
	pkg   string
	code  int
	msg   string // pure error msg
}

// New returns a new errorGroup whose error is nil.
func New(pkg string, code int) *errorGroup {
	return &errorGroup{
		pkg:  pkg,
		code: code,
	}
}

// Error returns error string
func (errGroup *errorGroup) Error() string {
	if errGroup.cause == nil {
		if errGroup.msg == "" {
			return fmt.Sprintf("ErrorGroup[Pkg=%s,Code=%d]", errGroup.pkg, errGroup.code)
		}

		return fmt.Sprintf("ErrorGroup[Pkg=%s,Code=%d,Msg=%s]",
			errGroup.pkg,
			errGroup.code,
			errGroup.msg,
		)
	}
	return errGroup.cause.Error()
}

// WithMsg appends a message to an errorGroup whose cause is nil.
func (errGroup *errorGroup) WithMsg(msg string) *errorGroup {
	if errGroup.cause != nil {
		return errGroup
	}

	return &errorGroup{
		cause: errGroup.cause,
		pkg:   errGroup.pkg,
		code:  errGroup.code,
		msg:   msg,
	}
}

// Wrap wraps an error.
// Note: use WrapWithMsg as much as possible
func (errGroup *errorGroup) Wrap(err error) *errorGroup {
	return &errorGroup{
		cause: errors.Wrap(err, errGroup.Error()),
		pkg:   errGroup.pkg,
		code:  errGroup.code,
	}
}

// WrapWithMsg wraps an error with a message.
func (errGroup *errorGroup) WrapWithMsg(err error, msg string) *errorGroup {
	return errGroup.WithMsg(msg).Wrap(err)
}

// getErrorPkgCode gets pkg and code from an error.
func as(err error) (*errorGroup, bool) {
	errGroup, ok := err.(*errorGroup)
	return errGroup, ok
}

// Is checks if a given err is an specific errorGroup
func (errGroup *errorGroup) Is(err error) bool {
	if asErrGroup, ok := as(err); ok {
		return asErrGroup.pkg == errGroup.pkg && asErrGroup.code == errGroup.code
	}
	return false
}

// GetErrcode gets error code from an error.
func GetErrcode(err error) (int, bool) {
	errGroup, ok := as(err)
	return errGroup.code, ok
}

// GetErrmsg gets pure error message from an error.
func GetErrmsg(err error) string {
	errGroup, ok := as(err)
	if ok {
		return errGroup.msg
	}
	return err.Error()
}
