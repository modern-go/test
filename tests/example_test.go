package tests

import (
	"context"
	. "github.com/modern-go/test"
	. "github.com/modern-go/test/must"
	"testing"
)

func Test(t *testing.T) {
	t.Run("one should not equal two", Case(func(ctx context.Context) {
		Assert(1 == 2)
	}))
}

func Test_one_should_not_equal_two(t *testing.T) {
	if 1 != 2 {
		t.Fatal("1==2")
	}
}
