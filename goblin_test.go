package goblin

import (
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAsync(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	g.Describe("Async test", func() {
		g.It("Should fail when Fail is called immediately", func(done Done) {
			g.Fail("Failed")
		})
		g.It("Should fail when Fail is called", func(done Done) {
			go func() {
				time.Sleep(100 * time.Millisecond)
				g.Fail("foo is not bar")
			}()
		})

		g.It("Should fail if done receives a parameter", func(done Done) {
			go func() {
				time.Sleep(100 * time.Millisecond)
				done("Error")
			}()
		})

		g.It("Should pass when done is called", func(done Done) {
			go func() {
				time.Sleep(100 * time.Millisecond)
				done()
			}()
		})

		g.It("Should fail if done has been called multiple times", func(done Done) {
			go func() {
				time.Sleep(100 * time.Millisecond)
				done()
				done()
			}()
		})
	})

	if !fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestAsyncSequence(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)
	var sequence []string

	g.Describe("Async test", func() {

		g.BeforeEach(func() {
			sequence = append(sequence, "global_before_each")
		})

		g.AfterEach(func() {
			sequence = append(sequence, "global_after_each")
		})

		g.Describe("nested", func() {

			g.BeforeEach(func() {
				sequence = append(sequence, "local_before_each")
			})

			g.AfterEach(func() {
				sequence = append(sequence, "local_after_each")
			})

			g.It("Should fail when Fail is called", func(done Done) {
				go func() {
					time.Sleep(100 * time.Millisecond)
					sequence = append(sequence, "test")
					done()
				}()
			})
		})

	})

	expected := []string{
		"global_before_each",
		"local_before_each",
		"test",
		"local_after_each",
		"global_after_each",
	}
	if !reflect.DeepEqual(expected, sequence) {
		t.Fatalf("Failed, expected:\n%s\n\ngot: %s\n",
			strings.Join(expected, "\n"),
			strings.Join(sequence, "\n"),
		)
	}
}

func TestAddNumbersSucceed(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {
		g.It("Should add numbers", func() {
			sum := 1 + 1
			g.Assert(sum).Equal(2)
		})
	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestAddNumbersFails(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {
		g.It("Should add numbers", func() {
			sum := 1 + 1
			g.Assert(sum).Equal(4)
		})
	})

	if !fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestMultipleIts(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	count := 0
	g.Describe("Numbers", func() {
		g.It("Should add numbers", func() {
			count++
			sum := 1 + 1
			g.Assert(sum).Equal(2)
		})

		g.It("Should add numbers", func() {
			count++
			sum := 1 + 1
			g.Assert(sum).Equal(4)
		})
	})

	if count != 2 {
		t.Fatal("Failed")
	}
}

func TestMultipleDescribes(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	count := 0
	g.Describe("Numbers", func() {

		g.Describe("Addition", func() {
			g.It("Should add numbers", func() {
				count++
				sum := 1 + 1
				g.Assert(sum).Equal(2)
			})
		})

		g.Describe("Subtraction", func() {
			g.It("Should subtract numbers", func() {
				count++
				sub := 5 - 6
				g.Assert(sub).Equal(1)
			})
		})
	})

	if count != 2 {
		t.Fatal("Failed")
	}
}

func TestPending(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {

		g.It("Should add numbers")

		g.Describe("Subtraction", func() {
			g.It("Should subtract numbers")
		})

	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestExcluded(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	count := 0
	g.Describe("Numbers", func() {

		g.Xit("Should add numbers", func() {
			count++
			sum := 1 + 1
			g.Assert(sum).Equal(2)
		})

		g.Describe("Subtraction", func() {
			g.Xit("Should subtract numbers", func() {
				count++
				sub := 5 - 6
				g.Assert(sub).Equal(1)
			})
		})

	})

	if count != 0 {
		t.Fatal("Failed")
	}

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestJustBeforeEach(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)
	const (
		before = iota
		beforeEach
		nBeforeEach
		justBeforeEach
		nJustBeforeEach
		it
		nIt
	)

	var (
		res [9]int
		i   int
	)

	g.Describe("Outer", func() {
		g.Before(func() {
			res[i] = before
			i++
		})

		g.BeforeEach(func() {
			res[i] = beforeEach
			i++
		})

		g.JustBeforeEach(func() {
			res[i] = justBeforeEach
			i++
		})

		g.It("should run all before handles by now", func() {
			res[i] = it
			i++
		})

		g.Describe("Nested", func() {
			g.BeforeEach(func() {
				res[i] = nBeforeEach
				i++
			})

			g.JustBeforeEach(func() {
				res[i] = nJustBeforeEach
				i++
			})

			g.It("should run all before handles by now", func() {
				res[i] = nIt
				i++
			})
		})
	})

	expected := [...]int{
		before,
		beforeEach,
		justBeforeEach,
		it,
		beforeEach,
		nBeforeEach,
		justBeforeEach,
		nJustBeforeEach,
		nIt,
	}

	if res != expected {
		t.Fatalf("expected %v to equal %v", res, expected)
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

		g.JustBeforeEach(func() {
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

			g.JustBeforeEach(func() {
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
		t.Fatal("Failed")
	}
}

func TestFailOnError(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {
		g.It("Does something", func() {
			g.Fail("Something")
		})
	})

	g.Describe("Errors", func() {
		g.It("Should fail with structs", func() {
			var s struct{ error string }
			s.error = "Error"
			g.Fail(s)
		})
	})

	if !fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestFailfOnError(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	g.Describe("Numbers", func() {
		g.It("Does something", func() {
			g.Failf("Something goes %s", "wrong")
		})
	})

	if !fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestRegex(t *testing.T) {
	fakeTest := testing.T{}
	os.Args = append(os.Args, "-goblin.run=matches")
	parseFlags()
	g := Goblin(&fakeTest)

	g.Describe("Test", func() {
		g.It("Doesn't match regex", func() {
			g.Fail("Regex shouldn't match")
		})

		g.It("It matches regex", func() {})
		g.It("It also matches", func() {})
	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}

	// Reset the regex so other tests can run
	runRegex = nil
}

func TestFailImmediately(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)
	reached := false
	g.Describe("Errors", func() {
		g.It("Should fail immediately for sync test", func() {
			g.Assert(false).IsTrue()
			reached = true
			g.Assert("foo").Equal("bar")
		})
		g.It("Should fail immediately for async test", func(done Done) {
			go func() {
				g.Assert(false).IsTrue()
				reached = true
				g.Assert("foo").Equal("bar")
				done()
			}()
		})
	})

	if reached {
		t.Fatal("Failed")
	}
}

func TestTimeout(t *testing.T) {
	fakeTest := testing.T{}
	os.Args = append(os.Args, "-goblin.timeout=10ms", "-goblin.run=")
	parseFlags()
	g := Goblin(&fakeTest)

	g.Describe("Test", func() {
		g.It("Should fail if test exceeds the specified timeout with sync test", func() {
			time.Sleep(100 * time.Millisecond)
		})

		g.It("Should fail if test exceeds the specified timeout with async test", func(done Done) {
			time.Sleep(100 * time.Millisecond)
			done()
		})
	})

	if !fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestItTimeout(t *testing.T) {
	fakeTest := testing.T{}
	os.Args = append(os.Args, "-goblin.timeout=10ms")
	parseFlags()
	g := Goblin(&fakeTest)

	g.Describe("Test", func() {
		g.It("Should override default timeout", func() {
			g.Timeout(20 * time.Millisecond)
			time.Sleep(15 * time.Millisecond)
		})

		g.It("Should revert for different it", func() {
			g.Assert(g.timeout).Equal(10 * time.Millisecond)
		})

	})
	if fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestIsNilAndIsNotNil(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	g.Describe("Test for IsNil", func() {
		g.It("Should assert successfully with nil value", func() {
			g.Assert(nil).IsNil()
		})
	})

	g.Describe("Test for IsNotNil", func() {
		g.It("Should assert successfully with not nil value", func() {
			g.Assert(struct{}{}).IsNotNil()
		})
	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}

	g.Describe("Test for IsNil with failed assertion", func() {
		g.It("Should fail", func() {
			g.Assert(100).IsNil()
		})
	})

	g.Describe("Test for IsNotNil with failed assertion", func() {
		g.It("Should fail", func() {
			g.Assert(nil).IsNotNil()
		})
	})

	if !fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestIsZeroAndIsNotZero(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	g.Describe("Test for IsZero", func() {
		g.It("Should assert successfully with int zero value", func() {
			g.Assert(0).IsZero()
		})

		g.It("Should assert successfully with float zero value", func() {
			g.Assert(0.0).IsZero()
		})

		g.It("Should assert successfully with string zero value", func() {
			g.Assert("").IsZero()
		})

		g.It("Should assert successfully with struct zero value", func() {
			g.Assert(struct{}{}).IsZero()
		})

		g.It("Should assert successfully with struct field with zero value", func() {
			g.Assert(struct{ value int }{value: 0}).IsZero()
		})
	})

	g.Describe("Test for IsNotZero", func() {
		g.It("Should assert successfully with int not zero value", func() {
			g.Assert(1).IsNotZero()
		})

		g.It("Should assert successfully with float  not zero value", func() {
			g.Assert(0.5).IsNotZero()
		})

		g.It("Should assert successfully with string not zero value", func() {
			g.Assert("ABC").IsNotZero()
		})

		g.It("Should assert successfully with struct not zero value", func() {
			g.Assert(struct{ value int }{value: 1}).IsNotZero()
		})
	})

	if fakeTest.Failed() {
		t.Fatal("Failed")
	}

	g.Describe("Test for IsZero with failed assertion", func() {
		g.It("Should fail", func() {
			g.Assert(100).IsZero()
		})

		g.It("Should fail", func() {
			g.Assert(1.0).IsZero()
		})

		g.It("Should fail", func() {
			g.Assert("A").IsZero()
		})

		g.It("Should fail", func() {
			g.Assert(struct{ value int }{value: 1}).IsZero()
		})
	})

	g.Describe("Test for IsNotZero with failed assertion", func() {
		g.It("Should fail", func() {
			g.Assert(0).IsNotZero()
		})

		g.It("Should fail", func() {
			g.Assert(0.0).IsNotZero()
		})

		g.It("Should fail", func() {
			g.Assert("").IsNotZero()
		})

		g.It("Should fail", func() {
			g.Assert(struct{}{}).IsNotZero()
		})
	})

	if !fakeTest.Failed() {
		t.Fatal("Failed")
	}
}

func TestSkip(t *testing.T) {
	fakeTest := testing.T{}

	g := Goblin(&fakeTest)

	count := 0
	g.Describe("Numbers", func() {

		g.Describe("Addition", func() {
			g.Skip("Should add numbers", func() {
				count |= 2
				sum := 1 + 1
				g.Assert(sum).Equal(2)
			})
		})

		g.Describe("Subtraction", func() {
			// This will skip all the following tests within this Describe block
			g.Skip()
			g.It("Should subtract numbers", func() {
				count |= 4
				sub := 5 - 6
				g.Assert(sub).Equal(1)
			})
		})

		g.Describe("Other", func() {
			g.It("Should not skip this since we're in a new block", func() {
				count |= 8
				check := !true
				g.Assert(check).Equal(false)
			})
		})

		// This should skip the following
		g.Skip()
		g.Describe("Nested skipping", func() {
			g.Describe("Should pass skipping down the chain", func() {
				g.It("Should skip tests that are nested", func() {
					count |= 16
					check := !false
					g.Assert(check).Equal(false)
				})
			})
		})

		g.Resume()
		// This should allow the next test to run
		g.It("Should run this since we resumed", func() {
			count |= 32
			g.Assert(count).IsNotZero()
		})

		// This should force the next test to skip
		g.SkipIf(func() bool {
			return true
		})
		g.It("Should skip this because our func returns true", func() {
			count |= 64
			g.Assert(count).IsZero()
		})
	})

	if fakeTest.Failed() {
		t.Fatal("Failed: suite failed")
	}
	if count != 8|32 {
		t.Fatal("Failed: incorrect count")
	}
}
