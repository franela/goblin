package goblin

import (
    "strings"
    "fmt"
    "strconv"
)
type Reporter interface {
    beginDescribe(string)
    endDescribe()
    begin()
    end()
    itFailed(string)
    itPassed(string)
}

type DetailedReporter struct {
    level, failed, passed int
}

func red(text string) string {
    return "\033[31m" + text + "\033[0m"
}

func gray(text string) string {
    return "\033[90m" + text + "\033[0m"
}

func (r *DetailedReporter) getSpace() (string) {
    return strings.Repeat(" ", (r.level+1)*2)
}

func (r *DetailedReporter) print(text string) {
    fmt.Printf("%v%v\n", r.getSpace(), text)
}

func (r *DetailedReporter) printWithCheck(text string) {
    fmt.Printf("%v\033[32m\u2713\033[0m %v\n", r.getSpace(), text)
}

func (r *DetailedReporter) beginDescribe(name string) {
    fmt.Println("")
    r.print(name)
    r.level++
}

func (r *DetailedReporter) endDescribe() {
    r.level--
}

func (r *DetailedReporter) itFailed(name string) {
    r.failed++
    r.print(red(strconv.Itoa(r.failed)+") "+name))
}

func (r *DetailedReporter) itPassed(name string) {
    r.passed++
    r.printWithCheck(gray(name))
}

func (r *DetailedReporter) begin() {
}

func (r *DetailedReporter) end() {
    fmt.Printf("\n\n \033[32m%d tests complete\033[0m\n \033[31m%d tests failed\033[0m\n\n", r.passed, r.failed)
}
