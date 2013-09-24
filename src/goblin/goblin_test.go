package goblin

import (
  "testing"
)

func TestAddNumbersSucceed(t *testing.T) {
  fakeTest := testing.T{}

  Describe("Numbers", func(d *D) {
    d.It("Should add numbers", func(t *T) {
      sum := 1+1
      t.Assert(sum).Equals(2)
    })
  })

  Goblin(&fakeTest)


  if fakeTest.Failed() {
    t.Fatal()
  }
}

func TestAddNumbersFails(t *testing.T) {
  fakeTest := testing.T{}


  Describe("Numbers", func(d *D) {
    d.It("Should add numbers", func(t *T) {
      sum := 1+1
      t.Assert(sum).Equals(4)
    })
  })

  Goblin(&fakeTest)

  if !fakeTest.Failed() {
    t.Fatal()
  }
}



func TestMultipleIts(t *testing.T) {
  fakeTest := testing.T{}

  Describe("Numbers", func(d *D) {

    d.It("Should add numbers", func(t *T) {
      sum := 1+1
      t.Assert(sum).Equals(4)
    })

    d.It("Should add numbers", func(t *T) {
      sum := 1+1
      t.Assert(sum).Equals(2)
    })
  })

  Goblin(&fakeTest)

  if !fakeTest.Failed() {
    t.Fatal()
  }
}



//func TestMultipleDescribes(t *testing.T) {
  //fakeTest := testing.T{}

  //Describe("Numbers", func(d *D) {

    //d.It("Should add numbers", func(t *T) {
      //sum := 1+1
      //t.Assert(sum).Equals(2)
    //})

    //d.Describe("Substraction", func(d *D) {
        //d.It("Should substract numbers ", func(t *T) {
          //sub := 5-5
          //t.Assert(sub).Equals(1)
        //})
    //})
  //})

  //Goblin(&fakeTest)

  //if !fakeTest.Failed() {
    //t.Fatal()
  //}
//}
