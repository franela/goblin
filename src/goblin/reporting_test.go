package goblin

import (
	"testing"
	"reflect"
)

type FakeReporter struct {
	describes []string
	fails []string
	passes []string
	ends int
    beginFlag, endFlag bool
}

func (r *FakeReporter) beginDescribe(name string) {
	r.describes = append(r.describes, name)
}

func (r *FakeReporter) endDescribe() {
	r.ends++
}

func (r *FakeReporter) itFailed(name string) {
	r.fails = append(r.fails, name)
}

func (r *FakeReporter) itPassed(name string) {
	r.passes = append(r.passes, name)
}

func (r *FakeReporter) begin() {
    r.beginFlag = true
}

func (r *FakeReporter) end() {
    r.endFlag = true
}

func TestReporting(t *testing.T) {
	fakeTest := &testing.T{}
	reporter := FakeReporter{}
	fakeReporter := Reporter(&reporter)	

	g := Goblin(fakeTest)
	g.SetReporter(fakeReporter)	

	g.Describe("One", func() {
		g.It("Foo", func() {
			g.Assert(0).Equals(1)
		})
		g.Describe("Two", func() {
			g.It("Bar", func() {
				g.Assert(0).Equals(0)
			})
		})
	})


	if !reflect.DeepEqual(reporter.describes, []string{"One", "Two"}) {
		t.FailNow()
	}
	if !reflect.DeepEqual(reporter.fails, []string{"Foo"}) {
		t.FailNow()
	}
	if !reflect.DeepEqual(reporter.passes, []string{"Bar"}) {
		t.FailNow()
	}
	if reporter.ends != 2 {
		t.FailNow()
	}
    if !reporter.beginFlag || !reporter.endFlag {
      t.FailNow()
    }
}
