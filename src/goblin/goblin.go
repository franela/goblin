package goblin

import (
  "testing"
  "fmt"
)

var parentDescribe Describe

type Runnable interface {
  Run() (bool)
}

type It struct {
  name string
  h func(*T)
  t *T
}

func (it It) Run() (bool) {
  fmt.Print(it.name)
  it.h(it.t)
  it.t.ran = true

  return !it.t.Failed()
}

type Describe struct {
  name string
  children []Runnable
}

func NewDescribe(name string, h func(*Describe)) {
  parentDescribe = Describe{name: name}
  h(&parentDescribe)
}

func (d *Describe) Describe(name string, h func(*Describe)) {
  describe := Describe{name: name}
  d.children = append(d.children, Runnable(describe))
  h(&describe)
}

func (d *Describe) It(name string, h func(t *T)) {
  it := It{name: name, h: h, t: &T{}}
  d.children = append(d.children, Runnable(it))
}

func (d Describe) Run() (bool) {
  fmt.Print(d.name)
  //TODO: run beforeEach
  succeed := true
  for _, r := range d.children {
    if !r.Run() {
      succeed = false
    }
  }
  //TODO: run afterEach
  return succeed
}

type T struct {
  testing.T

  ran bool
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

func Goblin(t *testing.T) {
  succeed := parentDescribe.Run()

  if !succeed {
    t.Fail()
  }
}
