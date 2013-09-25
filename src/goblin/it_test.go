package goblin

import (
    "testing"
)

func TestBeforeEach(t *testing.T) {
    fakeTest := testing.T{}


    Describe("Numbers", func(d *D) {
        before := 0

        d.BeforeEach(func() {
            before++
        })

        d.It("Should have called beforeEach", func(t *T) {
            t.Assert(before).Equals(1)
        })

        d.It("Should have called beforeEach also for this one", func(t *T) {
            t.Assert(before).Equals(2)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestMultipleBeforeEach(t *testing.T) {
    fakeTest := testing.T{}


    Describe("Numbers", func(d *D) {
        before := 0

        d.BeforeEach(func() {
            before++
        })

        d.BeforeEach(func() {
            before++
        })

        d.It("Should have called all the registered beforeEach", func(t *T) {
            t.Assert(before).Equals(2)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestNestedBeforeEach(t *testing.T) {
    fakeTest := testing.T{}


    Describe("Numbers", func(d *D) {
        before := 0

        d.BeforeEach(func() {
            before++
        })

        d.Describe("Addition", func(d *D) {
            d.BeforeEach(func() {
                before++
            })

            d.It("Should have called all the registered beforeEach", func(t *T) {
                t.Assert(before).Equals(2)
            })

            d.It("Should have called all the registered beforeEach also for this one", func(t *T) {
                t.Assert(before).Equals(4)
            })
        })

    })

    Goblin(&fakeTest)

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestAfterEach(t *testing.T) {
    fakeTest := testing.T{}

    after := 0
    Describe("Numbers", func(d *D) {

        d.AfterEach(func() {
            after++
        })

        d.It("Should call afterEach after this test", func(t *T) {
            t.Assert(after).Equals(0)
        })

        d.It("Should have called afterEach before this test ", func(t *T) {
            t.Assert(after).Equals(1)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() || after != 2 {
        t.Fatal()
    }
}

func TestMultipleAfterEach(t *testing.T) {
    fakeTest := testing.T{}


    after := 0
    Describe("Numbers", func(d *D) {

        d.AfterEach(func() {
            after++
        })

        d.AfterEach(func() {
            after++
        })

        d.It("Should call all the registered afterEach", func(t *T) {
            t.Assert(after).Equals(0)
        })
    })

    Goblin(&fakeTest)

    if fakeTest.Failed() || after != 2 {
        t.Fatal()
    }
}

func TestNestedAfterEach(t *testing.T) {
    fakeTest := testing.T{}

    after := 0
    Describe("Numbers", func(d *D) {

        d.AfterEach(func() {
            after++
        })

        d.Describe("Addition", func(d *D) {
            d.AfterEach(func() {
                after++
            })

            d.It("Should call all the registered afterEach", func(t *T) {
                t.Assert(after).Equals(0)
            })

            d.It("Should have called all the registered aftearEach", func(t *T) {
                t.Assert(after).Equals(2)
            })
        })

    })

    Goblin(&fakeTest)

    if fakeTest.Failed() || after != 4 {
        t.Fatal()
    }
}

