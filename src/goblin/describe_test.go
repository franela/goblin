package goblin

import (
    "testing"
)

func TestBefore(t *testing.T) {
    fakeTest := testing.T{}

    Describe("Numbers", func(d *D) {
        before := 0

        d.Before(func() {
            before++
        })

        d.It("Should have called before", func(t *T) {
            t.Assert(before).Equals(1)
        })

        d.It("Should have called before only once", func(t *T) {
            t.Assert(before).Equals(1)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestMultipleBefore(t *testing.T) {
    fakeTest := testing.T{}


    Describe("Numbers", func(d *D) {
        before := 0

        d.Before(func() {
            before++
        })

        d.Before(func() {
            before++
        })

        d.It("Should have called all the registered before", func(t *T) {
            t.Assert(before).Equals(2)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestNestedBefore(t *testing.T) {
    fakeTest := testing.T{}

    Describe("Numbers", func(d *D) {
        before := 0

        d.Before(func() {
            before++
        })

        d.Describe("Addition", func(d *D) {
            d.Before(func() {
                before++
            })

            d.It("Should have called all the registered before", func(t *T) {
                t.Assert(before).Equals(2)
            })

            d.It("Should have called all the registered before only once", func(t *T) {
                t.Assert(before).Equals(2)
            })
        })

    })

    Goblin(&fakeTest)

    if fakeTest.Failed() {
        t.Fatal()
    }
}


func TestAfter(t *testing.T) {
    fakeTest := testing.T{}

    after := 0
    Describe("Numbers", func(d *D) {

        d.After(func() {
            after++
        })

        d.It("Should call after only once", func(t *T) {
            t.Assert(after).Equals(0)
        })

        d.It("Should call after only once", func(t *T) {
            t.Assert(after).Equals(0)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() || after != 1 {
        t.Fatal()
    }
}

func TestMultipleAfter(t *testing.T) {
    fakeTest := testing.T{}


    after := 0
    Describe("Numbers", func(d *D) {

        d.After(func() {
            after++
        })

        d.After(func() {
            after++
        })

        d.It("Should call all the registered after", func(t *T) {
            t.Assert(after).Equals(0)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() && after != 2 {
        t.Fatal()
    }
}

func TestNestedAfter(t *testing.T) {
    fakeTest := testing.T{}

    after := 0
    Describe("Numbers", func(d *D) {

        d.After(func() {
            after++
        })

        d.Describe("Addition", func(d *D) {
            d.After(func() {
                after++
            })

            d.It("Should call all the registered after", func(t *T) {
                t.Assert(after).Equals(0)
            })

            d.It("Should have called all the registered after only once", func(t *T) {
                t.Assert(after).Equals(0)
            })
        })

    })

    Goblin(&fakeTest)

    if fakeTest.Failed() || after != 2 {
        t.Fatal()
    }
}
