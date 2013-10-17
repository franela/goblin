package goblin

import (
    "testing"
)

func TestAddNumbersSucceed(t *testing.T) {
    fakeTest := testing.T{}
    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {
        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equal(2)
        })
    })

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestAddNumbersFails(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {
        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equal(4)
        })
    })


    if !fakeTest.Failed() {
        t.Fatal()
    }
}



func TestMultipleIts(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {
        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equal(4)
        })

        g.It("Should add numbers", func() {
            sum := 1+1
            g.Assert(sum).Equal(2)
        })
    })


    if !fakeTest.Failed() {
        t.Fatal()
    }
}



func TestMultipleDescribes(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {

        g.Describe("Addition", func() {
           g.It("Should add numbers", func() {
                sum := 1+1
                g.Assert(sum).Equal(2)
            })
        })

        g.Describe("Substraction", func() {
            g.It("Should substract numbers ", func() {
                sub := 5-5
                g.Assert(sub).Equal(1)
            })
        })
    })


    if !fakeTest.Failed() {
        t.Fatal()
    }
}

func TestPending(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {

        g.It("Should add numbers")

        g.Describe("Substraction", func() {
           g.It("Should substract numbers")
        })

    })

    if fakeTest.Failed() {
        t.Fatal()
    }
}

func TestNotRunBeforesOrAfters(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)
    var count int

    g.Describe("Numbers", func() {
        g.Before(func() {
            count++
        })
        g.BeforeEach(func() {
            count++
        })

        g.After(func() {
            count++
        })
        g.AfterEach(func() {
            count++
        })

        g.Describe("Letters", func() {
            g.Before(func() {
                count++
            })
            g.BeforeEach(func() {
                count++
            })

            g.After(func() {
                count++
            })
            g.AfterEach(func() {
                count++
            })
        })
    })

    if count != 0 {
        t.Fatal()
    }
}

func TestFailOnError(t *testing.T) {
    fakeTest := testing.T{}

    g := Goblin(&fakeTest)

    g.Describe("Numbers", func() {
        g.It("Does something", func() {
            panic("Something")
        })
    })

    if !fakeTest.Failed() {
        t.Fatal()
    }
}
