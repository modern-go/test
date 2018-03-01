package tests

import (
	"context"
	"errors"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"github.com/modern-go/test/should"
	"os"
	"testing"
)

func Test_case(t *testing.T) {
	t.Run("should.Pass will not exit", test.Case(func(ctx context.Context) {
		should.Pass(1 == 1)
		should.Pass(1 == 2)
	}))
	t.Run("must.Pass will exit", test.Case(func(ctx context.Context) {
		must.Pass(1 == 1)
		must.Pass(1 == 2)
	}))
	t.Run("multiline", test.Case(func(ctx context.Context) {
		var f = func(i int) int { return i }
		must.Pass(struct{ i int }{
			f(100),
		} == struct{ i int }{
			f(101),
		})
	}))
	t.Run("attach details to assert", test.Case(func(ctx context.Context) {
		a := 1
		b := 2
		should.Pass(a > b, "a", a, "b", b)
	}))
	t.Run("equal", test.Case(func(ctx context.Context) {
		map1 := map[string]string{
			"a": "b",
			"c": "hello",
		}
		map2 := map[string]string{
			"a": "b",
			"c": "hi",
		}
		should.Equal(map1, map2)
	}))
	t.Run("nil", test.Case(func(ctx context.Context) {
		should.Nil(errors.New("hello"))
	}))
	t.Run("failed call", test.Case(func(ctx context.Context) {
		should.Call(os.Stat, "/tmp/no such file")
	}))
	t.Run("successful call", test.Case(func(ctx context.Context) {
		var stat os.FileInfo
		should.Call(os.Stat, "/bin/bash").Set(&stat)
		should.Equal("bash", stat.Name())
	}))
	t.Run("successful call 2", test.Case(func(ctx context.Context) {
		ret := should.Call(os.Stat, "/bin/bash")
		should.Equal("bash", ret[0].(os.FileInfo).Name())
	}))
	t.Run("json equal", test.Case(func(ctx context.Context) {
		must.JsonEqual(`{"a":{"b":"c"}}`, map[string]interface{}{
			"a": map[string]interface{}{
				"b": "c",
			},
		})
	}))
}
