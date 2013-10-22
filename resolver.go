package goblin

import (
    "runtime/debug"
    "strings"
)

func ResolveStack() ([]string) {
    return cleanStack(debug.Stack())
}


func cleanStack(stack []byte) []string {
    arrayStack := strings.Split(string(stack), "\n")
    var finalStack []string
    for i:=3; i<len(arrayStack); i++ {
        if strings.Contains(arrayStack[i], ".go") {
            finalStack = append(finalStack, arrayStack[i])
        }
    }
    return finalStack
}
