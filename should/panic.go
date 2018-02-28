package should

import (
	"github.com/modern-go/test"
	"github.com/modern-go/test/testify/assert"
	"runtime"
)

func Panic(f func()) (recovered interface{}) {
	t := test.CurrentT()
	test.Helper()
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Error("check failed")
		return
	}
	defer func() {
		recovered = recover()
		if assert.NotNil(t, recovered) {
			return
		}
		t.Error(test.ExtractFailedLines(file, line))
	}()
	f()
	return nil
}
