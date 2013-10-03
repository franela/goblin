package goblin

import (
  "testing"
  "time"
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
        g.reporter.begin()
        if d.run(g) {
            g.t.Fail()
        }
        g.reporter.end()
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
    g.reporter.beginDescribe(d.name)

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

    g.reporter.endDescribe()

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

    if it.h == nil {
        g.reporter.itIsPending(it.name)
        return true
    }
    //TODO: should handle errors for beforeEach
    it.parent.runBeforeEach()

    runIt(g, it.h)

    it.parent.runAfterEach()

    if it.failed {
        g.reporter.itFailed(it.name)
    } else {
        g.reporter.itPassed(it.name)
    }
    return it.failed

}

func Goblin (t *testing.T) (*G) {
    g := &G{t: t}
    g.reporter = Reporter(&DetailedReporter{})
    return g
}


func runIt (g *G, h func()) {
    defer timeTrack(time.Now(), g)
    h()
}


type G struct {
    t *testing.T
    parent *Describe
    currentIt *It
    reporter Reporter
}

func (g *G) SetReporter(r Reporter) {
    g.reporter = r
}

func (g *G) It(name string, h ...func()) {
    it := &It{name:name, parent:g.parent}
    if len(h) > 0 {
        it.h = h[0]
        g.parent.children = append(g.parent.children, Runnable(it))
    } else {
        g.parent.children = append(g.parent.children, Runnable(it))
    }
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


func timeTrack(start time.Time, g *G) {
        g.reporter.itTook(time.Since(start))
}
