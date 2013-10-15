package goblin

import (
    "testing"
)

var failed = false

func fakeFail(message string, callerSkip ...int) {
    failed = true;
}

func TestEqual(t *testing.T) {
    a := Assertion{src: 1, fail: fakeFail}
    a.Equal(1)

    if failed {
      t.FailNow()
    }

    a = Assertion{src: "foo", fail: fakeFail}
    a.Equal("foo")

    if failed {
      t.FailNow()
    }

    a = Assertion{src: map[string]string{"foo": "bar"}, fail: fakeFail}
    a.Equal(map[string]string{"foo": "bar"})

    if failed {
      t.FailNow()
    }
}
