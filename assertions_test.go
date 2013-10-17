package goblin

import (
    "testing"
)

var failed = false

func TestEqual(t *testing.T) {
    a := Assertion{src: 1}
    a.Equal(1)

    if failed {
      t.FailNow()
    }

    a = Assertion{src: "foo"}
    a.Equal("foo")

    if failed {
      t.FailNow()
    }

    a = Assertion{src: map[string]string{"foo": "bar"}}
    a.Equal(map[string]string{"foo": "bar"})

    if failed {
      t.FailNow()
    }
}
