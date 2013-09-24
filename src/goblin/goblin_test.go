package goblin

import (
  "testing"
)

func TestAddNumbersSucceed(t *testing.T) {
  ft := testing.T{}

  var test *T

  NewDescribe("Numbers", func(d *Describe) {
    d.It("Should add numbers", func(t *T) {
      test = t
      sum := 1+1
      t.Assert(sum).Equals(2)
    })
  })

  Goblin(&ft)

  if !test.ran {
    t.Fatal()
  }

  if ft.Failed() {
    t.Fatal()
  }
}

func TestAddNumbersFails(t *testing.T) {
  ft := testing.T{}

  var test *T

  NewDescribe("Numbers", func(d *Describe) {
    d.It("Should add numbers", func(t *T) {
      test = t
      sum := 1+1
      t.Assert(sum).Equals(4)
    })
  })

  Goblin(&ft)

  if !test.ran {
    t.Fatal()
  }

  if !ft.Failed() {
    t.Fatal()
  }
}
