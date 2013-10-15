package goblin

import (
        "runtime"
)

func ResolveCaller(depth int) (string, int) {
    _, file, line, _ := runtime.Caller(depth)
    return file, line
}
