package goblin

import (
	"testing"
)

func TestBefore(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {
		before := 0

		g.Before(func() {
			before++
		})

		g.It("Should have called before", func() {
			g.Assert(before).Equal(1)
		})

		g.It("Should have called before only once", func() {
			g.Assert(before).Equal(1)
		})
	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestMultipleBefore(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {
		before := 0

		g.Before(func() {
			before++
		})

		g.Before(func() {
			before++
		})

		g.It("Should have called all the registered before", func() {
			g.Assert(before).Equal(2)
		})
	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestNestedBefore(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {
		before := 0

		g.Before(func() {
			before++
		})

		g.Describe("Addition", func() {
			g.Before(func() {
				before++
			})

			g.It("Should have called all the registered before", func() {
				g.Assert(before).Equal(2)
			})

			g.It("Should have called all the registered before only once", func() {
				g.Assert(before).Equal(2)
			})
		})

	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestAfter(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)
	after := 0
	g.Describe("Numbers", func() {

		g.After(func() {
			after++
		})

		g.It("Should call after only once", func() {
			g.Assert(after).Equal(0)
		})

		g.It("Should call after only once", func() {
			g.Assert(after).Equal(0)
		})
	})

	if fakeTest.Failed() || after != 1 {
		t.Fatal("Failed")
	}
}

func TestMultipleAfter(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	after := 0
	g.Describe("Numbers", func() {

		g.After(func() {
			after++
		})

		g.After(func() {
			after++
		})

		g.It("Should call all the registered after", func() {
			g.Assert(after).Equal(0)
		})
	})

	if fakeTest.Failed() && after != 2 {
		t.Fatal("Failed")
	}
}

func TestNestedAfter(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)
	after := 0
	g.Describe("Numbers", func() {

		g.After(func() {
			after++
		})

		g.Describe("Addition", func() {
			g.After(func() {
				after++
			})

			g.It("Should call all the registered after", func() {
				g.Assert(after).Equal(0)
			})

			g.It("Should have called all the registered after only once", func() {
				g.Assert(after).Equal(0)
			})
		})

	})

	if fakeTest.Failed() || after != 2 {
		t.Fatal("Failed")
	}
}

func TestSkipDescribe(t *testing.T) {
	g := Goblin(t)

	g.Describe("Describe will run", func() {

		g.It("This will run", func() {
			g.Assert(4).Equal(4)
		})
	})

	g.Skip.Describe("Describe will not run", func() {

		g.Before(func() {
			t.Fatal("This Before() should not run")
		})
		g.After(func() {
			t.Fatal("This After() should not run")
		})

		g.BeforeEach(func() {
			t.Fatal("This BeforeEach() should not run")
		})
		g.AfterEach(func() {
			t.Fatal("This AfterEach() should not run")
		})

		g.JustBeforeEach(func() {
			t.Fatal("This JustBeforeEach() should not run")
		})

		g.It("This will not run", func() {
			t.Fatal("Failed")
		})

		g.Describe("Describe will not run also", func() {
			g.Before(func() {
				t.Fatal("This Before() should not run")
			})
			g.After(func() {
				t.Fatal("This After() should not run")
			})
			g.It("This will not run also", func() {
				t.Fatal("Failed")
			})
		})
	})

	g.Describe("Last describe will run", func() {

		counter := 0

		g.Before(func() {
			counter++
		})
		g.BeforeEach(func() {
			counter++
		})
		g.JustBeforeEach(func() {
			counter++
		})
		g.It("This will run", func() {
			g.Assert(counter).Equal(3)
		})
	})
}
