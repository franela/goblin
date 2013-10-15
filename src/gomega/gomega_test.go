package gomega

import (
    "testing"
    . "github.com/onsi/gomega"
    . "goblin"
)

func TestGoMegaIntegration(t *testing.T) {
    fakeTest := testing.T{}


    g := Goblin(&fakeTest, )
    RegisterFailHandler(g.Fail)

    g.Describe("Numbers", func() {
        g.It("Should add numbers (pass)", func() {
            sum := 1+1
            Expect(sum).To(Equal(2))
        })
        g.It("Should add numbers (fail)", func() {
            sum := 1+1
            Expect(sum).To(Equal(4))
        })
    })


    if !fakeTest.Failed() {
        t.Fatal()
    }
}
