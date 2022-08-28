package errorgrouppkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_Error(t *testing.T) {
	err := New("pkg", 1)
	expectedErrMsg := "ErrorGroup[Pkg=pkg,Code=1]"
	assert.Equal(t, expectedErrMsg, err.Error())
}

func Test_WithMsg(t *testing.T) {
	err := New("pkg", 1)
	errWithMsg := err.WithMsg("query_db_failed")
	assert.Equal(t, "ErrorGroup[Pkg=pkg,Code=1][Msg=query_db_failed]", errWithMsg.Error())
}

func Test_Wrap(t *testing.T) {
	err := fmt.Errorf("table not found")
	errWrap := New("pkg", 1).Wrap(err)
	assert.Equal(t, "ErrorGroup[Pkg=pkg,Code=1]: table not found", errWrap.Error())
}

func Test_WrapWithMsg(t *testing.T) {
	err := fmt.Errorf("table not found")
	errWrapWithMsg := New("pkg1", 1).WrapWithMsg(err, "query_db_failed")
	assert.Equal(t, "ErrorGroup[Pkg=pkg1,Code=1][Msg=query_db_failed]: table not found", errWrapWithMsg.Error())
}

func Test_Is(t *testing.T) {
	var err error = New("pkg", 1)
	var err2 = New("pkg", 1)
	assert.Equal(t, true, err2.Is(err))
}

func Test_GetErrcode(t *testing.T) {
	var err error
	err = New("pkg", 1)
	errcode, ok := GetErrcode(err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, errcode)
}
