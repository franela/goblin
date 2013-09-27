package goblin


type Assertion struct {
  src int
  it *It
}

func (a *Assertion) Equals(dst int) {
    if dst != a.src {
        a.it.failed = true
    }
}
