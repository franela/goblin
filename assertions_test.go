package goblin

import (
	"testing"
)

var failed = false

func TestEqual(t *testing.T) {
	a := Assertion{src: 1}
	a.Equal(1)
	a.Eql(1)

	if failed {
		t.FailNow()
	}

	a = Assertion{src: "foo"}
	a.Equal("foo")
	a.Eql("foo")

	if failed {
		t.FailNow()
	}

	a = Assertion{src: map[string]string{"foo": "bar"}}
	a.Equal(map[string]string{"foo": "bar"})
	a.Eql(map[string]string{"foo": "bar"})

	if failed {
		t.FailNow()
	}
}

func TestIsTrue(t *testing.T) {
	a := Assertion{src: true}
	a.IsTrue()

	if failed {
		t.FailNow()
	}
}

func TestIsFalse(t *testing.T) {
	a := Assertion{src: false}
	a.IsFalse()

	if failed {
		t.FailNow()
	}
}
