package goblin

import (
    "testing"
)


func TestEqual(t *testing.T) {
    fakeIt := &It{}

    a := Assertion{src: 1, it: fakeIt}
    a.Equal(1)

    if len(fakeIt.failures) != 0 {
      t.FailNow()
    }

    a = Assertion{src: "foo", it: fakeIt}
    a.Equal("foo")

    if len(fakeIt.failures) != 0 {
      t.FailNow()
    }

    a = Assertion{src: map[string]string{"foo": "bar"}, it: fakeIt}
    a.Equal(map[string]string{"foo": "bar"})

    if len(fakeIt.failures) != 0 {
      t.FailNow()
    }
}
