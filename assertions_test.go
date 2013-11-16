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

func TestBeTrue(t *testing.T) {
    a := Assertion{src: true}
    a.BeTrue()

    if failed {
      t.FailNow()
    }
}

func TestBeFalse(t *testing.T) {
    a := Assertion{src: false}
    a.BeFalse()

    if failed {
      t.FailNow()
    }
}
