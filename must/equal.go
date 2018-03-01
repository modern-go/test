package must

import (
	"github.com/modern-go/test"
	"github.com/stretchr/testify/assert"
	"runtime"
)

//go:noinline
func Equal(expected interface{}, actual interface{}) {
	t := test.CurrentT()
	if assert.Equal(t, expected, actual) {
		return
	}
	test.Helper()
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("check failed")
		return
	}
	t.Fatal(test.ExtractFailedLines(file, line))
}

//go:noinline
func AssertEqual(expected interface{}, actual interface{}) {
	t := test.CurrentT()
	if assert.Equal(t, expected, actual) {
		return
	}
	test.Helper()
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("check failed")
		return
	}
	t.Fatal(test.ExtractFailedLines(file, line))
}
