package goblin


import (
        "testing"
)

type T struct {
  testing.T
}

func (t *T) Assert(num int) (*Assertion) {
  return &Assertion{ t: t, src: num }
}

type Assertion struct {
  src int
  t *T
}

func (a *Assertion) Equals(dst int) {
  if dst != a.src {
    a.t.Fail()
  }
}
