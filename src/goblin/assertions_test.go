package goblin

import (
    "testing"
)


func TestEqual(t *testing.T) {
    fakeIt := &It{}

    a := Assertion{src: 1, it: fakeIt}
    a.Equal(1)

    if fakeIt.failed {
      t.FailNow()
    }

    a = Assertion{src: "foo", it: fakeIt}
    a.Equal("foo")

    if fakeIt.failed {
      t.FailNow()
    }

    a = Assertion{src: map[string]string{"foo": "bar"}, it: fakeIt}
    a.Equal(map[string]string{"foo": "bar"})

    if fakeIt.failed {
      t.FailNow()
    }
}
