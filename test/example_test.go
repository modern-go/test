package test

import (
	"context"
	. "github.com/modern-go/test"
	. "github.com/modern-go/test/must"
	"testing"
)

func Test(t *testing.T) {
	t.Run("1 != 2", Case(func(ctx context.Context) {
		Assert(1 == 2)
	}))
}
