package goblin

import (
	"sync"
	"testing"
	"time"
)

func TestG_It_AsyncDone_Race(t *testing.T) {
	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	g.Describe("Async test", func() {
		g.It("Should not create a data race if done() is called multiple times", func(done Done) {
			go func() {
				time.Sleep(100 * time.Millisecond)
				done()
				done()
			}()
		})
	})
}

func TestG_It_Fail_Race(t *testing.T) {
	g := Goblin(new(testing.T))

	g.Describe("Synchronous test", func() {
		g.It("Should not create a data race on fail", func() {
			g.Fail("Failed")
		})
	})
}

func TestG_It_Assert_Race(t *testing.T) {
	g := Goblin(new(testing.T))
	g.SetReporter(Reporter(new(FakeReporter)))

	g.Describe("Should not create a data race", func() {
		g.It("Should fail", func() {
			g.Assert(0).Equal(1)
		})
		g.It("Should pass", func() {
		})
	})
}

func TestG_Parallel_New_Goblin(t *testing.T) {
	wg := new(sync.WaitGroup)
	const cnt = 2

	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			Goblin(new(testing.T))
			wg.Done()
		}()
	}

	wg.Wait()
}
