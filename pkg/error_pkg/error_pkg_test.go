package errorpkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_Error(t *testing.T) {
	err := New("pkg", 1)
	assert.Equal(t, "Error{Pkg=pkg,Code=1}", err.Error())
}

func Test_WithMessage(t *testing.T) {
	err := New("pkg", 1)
	errWithMsg := err.WithMsg("query_db_failed")
	assert.Equal(t, "Error{Pkg=pkg,Code=1}query_db_failed", errWithMsg.Error())
}

func Test_Wrap(t *testing.T) {
	err := fmt.Errorf("table not found")
	errWrap := New("pkg2", 2).Wrap(err)
	assert.Equal(t, "Error{Pkg=pkg2,Code=2}: table not found", errWrap.Error())
}

func ()  {
	errUnwrap := errWrap.Unwrap()
	fmt.Printf("%s\n\n%v\n\n%+v\n", errUnwrap, errUnwrap, errUnwrap)
}
