package goblin

import (
        "testing"
        "os"
        "runtime"
)

func TestResolver(t *testing.T) {

    g := Goblin(t)

    g.Describe("Resolver", func() {

        g.It("Should resolve the caller filename ", func() {
            file, _:= ResolveCaller(1)
            cwd, _ := os.Getwd()
            g.Assert(file).Equal(cwd+"/resolver_test.go")
        })

        g.It("Should resolve the caller line ", func() {
            _, _, currentLine, _ := runtime.Caller(0)
            _, line:= ResolveCaller(1)
            g.Assert(line).Equal(currentLine+1)
        })

    })
}
