package errorgrouppkg

import (
	"fmt"

	"github.com/pkg/errors"
)

// errorGroup defines an error type with package and code
type errorGroup struct {
	cause error
	pkg   string
	code  int
}

// New returns a new errorGroup whose error is nil.
func New(pkg string, code int) *errorGroup {
	return &errorGroup{
		pkg:  pkg,
		code: code,
	}
}

const (
	errorGroupPkgCodeFormat = "ErrorGroup[Pkg=%s,Code=%d]"
	errorGroupWithMsgFormat = "%s[Msg=%s]"
)

// Error returns error string
func (errGroup *errorGroup) Error() string {
	if errGroup.cause == nil {
		return fmt.Sprintf(errorGroupPkgCodeFormat, errGroup.pkg, errGroup.code)
	}
	return errGroup.cause.Error()
}

// WithMsg appends a message to an errorGroup whose cause is nil.
func (errGroup *errorGroup) WithMsg(msg string) *errorGroup {
	if errGroup.cause != nil {
		return errGroup
	}

	return &errorGroup{
		cause: fmt.Errorf(errorGroupWithMsgFormat, errGroup.Error(), msg),
		pkg:   errGroup.pkg,
		code:  errGroup.code,
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
func getErrorPkgCode(err error) (string, int, bool) {
	errGroup, ok := err.(*errorGroup)
	return errGroup.pkg, errGroup.code, ok
}

// Is checks if a given err is an specific errorGroup
func (errGroup *errorGroup) Is(err error) bool {
	if pkg, code, ok := getErrorPkgCode(err); ok {
		return pkg == errGroup.pkg && code == errGroup.code
	}
	return false
}

// GetErrcode gets code from an error.
func GetErrcode(err error) (int, bool) {
	_, code, ok := getErrorPkgCode(err)
	return code, ok
}
