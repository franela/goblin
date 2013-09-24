package goblin

import (
  "testing"
  "fmt"
)

var parentDescribe *D

type Runnable interface {
  Run() (bool)
}

type It struct {
  name string
  h func(*T)
  t *T
}

func (it It) Run() (bool) {
  fmt.Println(it.name)
  it.h(it.t)

  return !it.t.Failed()
}

type D struct {
  name string
  children []Runnable
}

func Describe(name string, h func(*D)) {
  parentDescribe = &D{name: name}
  h(parentDescribe)
}

func (d *D) Describe(name string, h func(*D)) {
  describe := &D{name: name}
  d.addChild(Runnable(describe))
  h(describe)
  fmt.Println(d.children,"AFTER IT", describe, Runnable(describe))
}

func (d *D) addChild(r Runnable) {
  d.children = append(d.children, r)
}

func (d *D) It(name string, h func(t *T)) {
  it := It{name: name, h: h, t: &T{}}
  d.addChild(Runnable(it))
}

func (d D) Run() (bool) {
  fmt.Println(d.name)
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


func Goblin(t *testing.T) {
  fmt.Println(parentDescribe, "PARENT")
  succeed := parentDescribe.Run()

  if !succeed {
    t.Fail()
  }
}
