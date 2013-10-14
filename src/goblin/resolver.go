package goblin

import (
        "runtime"
)

func ResolveCaller() (string, int) {
    _, file, line, _ := runtime.Caller(1)
    return file, line
}

func ResolveStack() string {
    return ""
}
