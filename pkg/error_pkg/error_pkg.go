package errorpkg

import (
	"fmt"

	"github.com/pkg/errors"
)

// errorPackage defines an error type with package and code
type errorPackage struct {
	cause error
	pkg   string
	code  int
}

// New returns a new errorPackage whose error is nil.
func New(pkg string, code int) *errorPackage {
	return &errorPackage{
		pkg:  pkg,
		code: code,
	}
}

// Error returns error string
func (errPkg *errorPackage) Error() string {
	if errPkg.cause == nil {
		return fmt.Sprintf("Error{Pkg=%s,Code=%d}", errPkg.pkg, errPkg.code)
	}
	return errPkg.cause.Error()
}

// WithMsg appends a message to an errorPackage whose error is nil.
func (errPkg *errorPackage) WithMsg(msg string) *errorPackage {
	if errPkg.cause != nil {
		return errPkg
	}

	return &errorPackage{
		cause: fmt.Errorf("%s%s", errPkg.Error(), msg),
		pkg:   errPkg.pkg,
		code:  errPkg.code,
	}
}

// Wrap wraps an error.
// Note: I personlly recommend using WrapWithMsg as much as possible.
func (errPkg *errorPackage) Wrap(err error) *errorPackage {
	return &errorPackage{
		cause: errors.Wrap(err, errPkg.Error()),
		pkg:   errPkg.pkg,
		code:  errPkg.code,
	}
}

// WrapWithMsg wraps an error with a message.
func (errPkg *errorPackage) WrapWithMsg(err error, msg string) *errorPackage {
	return errPkg.WithMsg(msg).Wrap(err)
}

// Is checks if a given err is an errorPackeg
func (errPkg *errorPackage) Is(err error) bool {
	if pkg, code, ok := getErrorPackageCode(err); ok {
		return pkg == errPkg.pkg && code == errPkg.code
	}
	return false
}

// getErrorPackageCode gets pkg and code from an error.
func getErrorPackageCode(err error) (string, int, bool) {
	errPkg, ok := err.(*errorPackage)
	return errPkg.pkg, errPkg.code, ok
}

// GetErrorCode gets code from an error.
func GetErrorCode(err error) (int, bool) {
	_, code, ok := getErrorPackageCode(err)
	return code, ok
}
