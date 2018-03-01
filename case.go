package test

import (
	"context"
	"github.com/modern-go/gls"
	"reflect"
	"testing"
)

var testingTType = reflect.TypeOf((*testing.T)(nil))

func Case(testCase func(ctx context.Context)) func(t *testing.T) {
	return func(t *testing.T) {
		goid := gls.GoID()
		gls.ResetGls(goid, map[interface{}]interface{}{
			testingTType: t,
		})
		ctx := context.Background()
		defer func() {
			gls.DeleteGls(goid)
		}()
		testCase(ctx)
	}
}

func Skip(args ...interface{}) {
	CurrentT().Skip(args...)
}

func Skipf(format string, args ...interface{}) {
	CurrentT().Skipf(format, args...)
}

func CurrentT() *testing.T {
	t, found := gls.Get(testingTType).(*testing.T)
	if !found {
		panic("test not started with test.Case()")
	}
	return t
}
