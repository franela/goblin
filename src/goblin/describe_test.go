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
