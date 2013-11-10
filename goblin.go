package goblin

import (
    "testing"
    "time"
    "runtime"
    "fmt"
)

type Done func(error ...interface{})

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
    hasTests bool
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

    failed := ""

    if d.hasTests {
        for _, b := range d.befores {
            b()
        }
    }

    for _, r := range d.children {
        if r.run(g) {
            failed = "true"
        }
    }

    if d.hasTests {
        for _, a := range d.afters {
            a()
        }
    }

    g.reporter.endDescribe()

    return failed != ""
}

type Failure struct {
    stack []string
    testName string
    message string
}

type It struct {
    h interface{}
    name string
    parent *Describe
    failure *Failure
    reporter Reporter
    isAsync bool
}

func (it *It) run(g *G) (bool) {
    g.currentIt = it

    if it.h == nil {
        g.reporter.itIsPending(it.name)
        return false
    }
    //TODO: should handle errors for beforeEach
    it.parent.runBeforeEach()

    runIt(g, it.h)

    it.parent.runAfterEach()

    failed := false
    if it.failure != nil {
        failed = true
    }

    if failed {
        g.reporter.itFailed(it.name)
        g.reporter.failure(it.failure)
    } else {
        g.reporter.itPassed(it.name)
    }
    return failed
}

func (it *It) failed(msg string, stack []string) {
    it.failure = &Failure{stack:stack, message:msg, testName: it.parent.name + " " + it.name}
}

func Goblin (t *testing.T) (*G) {
    g := &G{t: t}
    g.reporter = Reporter(&DetailedReporter{})
    return g
}


func runIt (g *G, h interface{}) {
    defer timeTrack(time.Now(), g)

    // We do this to recover from panic, which is how we know that the test failed.
    /*
    defer func() {
        if r := recover(); r != nil {
            stack := ResolveStack()
            e := fmt.Sprintf("%v", r)
            g.currentIt.failed(e, stack)
        }
    }()
*/
    if call, ok := h.(func()); ok {
        // the test is synchronous
        call()
    } else if call, ok := h.(func(Done)); ok {
        g.currentIt.isAsync = true
        // the test is asynchronous
        g.shouldContinue = make(chan bool)
        doneCalled := 0
        go func() {call(func(msg ...interface{}) {
            if len(msg) > 0 {
                fmt.Println("lalal")
                g.Fail(msg)
            } else {
                doneCalled++
                if doneCalled > 1 {
                    g.Fail("Done called multiple times")
                }
                g.shouldContinue <- true
            }
        })} ()
        <- g.shouldContinue
    } else {
        panic("Not implemented.")
    }
}


type G struct {
    t *testing.T
    parent *Describe
    currentIt *It
    reporter Reporter
    shouldContinue chan bool
}

func (g *G) SetReporter(r Reporter) {
    g.reporter = r
}

func (g *G) It(name string, h ...interface{}) {
    it := &It{name:name, parent:g.parent, reporter: g.reporter}
    notifyParents(g.parent)
    if len(h) > 0 {
        it.h = h[0]
    }
    g.parent.children = append(g.parent.children, Runnable(it))
}

func notifyParents(d *Describe) {
    d.hasTests = true
    if d.parent != nil {
        notifyParents(d.parent)
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

func (g *G) Assert(src interface{}) (*Assertion) {
    return &Assertion{src: src, fail: g.Fail}
}


func timeTrack(start time.Time, g *G) {
    g.reporter.itTook(time.Since(start))
}

func (g *G) Fail(error interface{}) {
    stack := ResolveStack()
    message := fmt.Sprintf("%v", error)
    g.currentIt.failed(message, stack)
    if g.shouldContinue != nil {
        g.shouldContinue <- true
    }
    if g.currentIt.isAsync {
       runtime.Goexit()
    }
}
