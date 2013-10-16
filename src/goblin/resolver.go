package goblin

import (
    "runtime"
    "strings"
)

func ResolveCaller() (string, int) {
    var filename string
    var line int

    for depth:=0; !strings.HasSuffix(filename, "_test.go"); depth++ {
        _, filename, line, _ = runtime.Caller(depth)
    }
    return filename, line
}
