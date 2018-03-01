# test

[![Sourcegraph](https://sourcegraph.com/github.com/modern-go/test/-/badge.svg)](https://sourcegraph.com/github.com/modern-go/test?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/modern-go/test)
[![rcard](https://goreportcard.com/badge/github.com/modern-go/test)](https://goreportcard.com/report/github.com/modern-go/test)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/modern-go/test/master/LICENSE)

Make go unit test more readable

* C style assert, with clear error message
* Diff object, if unequal, tells exactly where is the difference
* JsonEqual to compare complex object graph

# Assert

```go
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
```

If test failed, the error message will tell you where it failed and what has failed

```
example_test.go:12: failed: Assert(1 == 2)
```

# Equal/NotEqual

the implementation of diff is copied from testify.

```go
map1 := map[string]string{
"a": "b",
"c": "hello",
}
map2 := map[string]string{
"a": "b",
"c": "hi",
}
must.Equal(map1, map2)
```

the diff will be printed

```
Error: Not equal:
expected: map[string]string{"a":"b", "c":"hello"}
actual : map[string]string{"a":"b", "c":"hi"}

Diff:
--- Expected
+++ Actual
@@ -2,3 +2,3 @@
(string) (len=1) "a": (string) (len=1) "b",
- (string) (len=1) "c": (string) (len=5) "hello"
+ (string) (len=1) "c": (string) (len=2) "hi"
}
case_test.go:39: failed: must.Equal(map1, map2)
```

# JsonEqual

check complex object graph using json

```go
must.JsonEqual(`{"a":{"b":"c"}}`, map[string]interface{}{
    "a": map[string]interface{}{
        "b": "c",
    },
})
```

# must v.s. should

* must: when assertion failed, the test case will skip remaining assertions, skip to next test case
* should: when assertion failed, the test case will continue check remaining assertions