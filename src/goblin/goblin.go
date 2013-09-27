package goblin

import (
  "testing"
  "fmt"
)

type Runnable interface {
    run(*G) (bool)
}




func (g *G) Describe(name string, h func()) {
    d := &Describe{name:name, h:h, parent:g.parent}

    if d.parent != nil {
        d.parent.children = append(d.parent.children, Runnable(d))
    }

    g.parent = d

    h()

    g.parent = d.parent

    if g.parent == nil {
        if d.run(g) {
            g.t.Fail()
        }
    }
}

type Describe struct {
    name string
    h func()
    children []Runnable
    befores []func()
    afters []func()
    afterEach []func()
    beforeEach []func()
    parent *Describe
}

func (d *Describe) runBeforeEach() {
    if d.parent != nil {
        d.parent.runBeforeEach()
    }

    for _, b := range d.beforeEach {
        b()
    }
}

func (d *Describe) runAfterEach() {

    if d.parent != nil {
      d.parent.runAfterEach()
    }

    for _, a := range d.afterEach {
      a()
    }
}


func (d *Describe) run(g *G) (bool) {
    failed := false

    for _, b := range d.befores {
        b()
    }

    for _, r := range d.children {
        if r.run(g) {
            failed = true
        }
    }

    for _, a := range d.afters {
        a()
    }

    return failed
}

type It struct {
    h func()
    name string
    parent *Describe
    failed bool
}

func (it *It) run(g *G) (bool) {
    g.currentIt = it
    //TODO: should handle errors for beforeEach
    it.parent.runBeforeEach()

    fmt.Println(it.name)

    it.h()

    it.parent.runAfterEach()

    return it.failed

}

func Goblin (t *testing.T) (*G) {
    g := &G{t: t}
    return g
}


type G struct {
    t *testing.T
    parent *Describe
    currentIt *It

}


func (g *G) It(name string, h func()) {
    it := &It{name:name, h:h, parent:g.parent}
    g.parent.children = append(g.parent.children, Runnable(it))
}

func (g *G) Before(h func()) {
    g.parent.befores = append(g.parent.befores, h)
}

func (g *G) BeforeEach(h func()) {
    g.parent.beforeEach = append(g.parent.beforeEach, h)
}

func (g *G) After(h func()) {
    g.parent.afters = append(g.parent.afters, h)
}

func (g *G) AfterEach(h func()) {
    g.parent.afterEach = append(g.parent.afterEach, h)
}

func (g *G) Assert(src int) (*Assertion) {
    return &Assertion{src: src , it: g.currentIt}
}

