package goblin

import (
    "strings"
    "fmt"
    "strconv"
    "time"
)
type Reporter interface {
    beginDescribe(string)
    endDescribe()
    begin()
    end()
    failure(*Failure)
    itTook(time.Duration)
    itFailed(string)
    itPassed(string)
    itIsPending(string)
}


type DetailedReporter struct {
    level, failed, passed, pending int
    failures []*Failure
    executionTime, totalExecutionTime time.Duration
}

func red(text string) string {
    return "\033[31m" + text + "\033[0m"
}

func gray(text string) string {
    return "\033[90m" + text + "\033[0m"
}

func cyan(text string) string {
    return "\033[36m" + text + "\033[0m"
}

func (r *DetailedReporter) getSpace() (string) {
    return strings.Repeat(" ", (r.level+1)*2)
}

func (r *DetailedReporter) failure(failure *Failure) {
    r.failures = append(r.failures, failure)
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

func (r *DetailedReporter) itTook(duration time.Duration) {
    r.executionTime = duration
    r.totalExecutionTime += duration
}

func (r *DetailedReporter) itFailed(name string) {
    r.failed++
    r.print(red(strconv.Itoa(r.failed)+") "+name))
}

func (r *DetailedReporter) itPassed(name string) {
    r.passed++
    r.printWithCheck(gray(name))
}

func (r *DetailedReporter) itIsPending(name string) {
    r.pending++
    r.print(cyan("- "+name))
}

func (r *DetailedReporter) begin() {
}

func (r *DetailedReporter) end() {
    fmt.Printf("\n\n \033[32m%d tests complete\033[0m \033[90m(%d ms)\033[0m\n", r.passed, r.totalExecutionTime / time.Millisecond)

    if r.pending > 0 {
        fmt.Printf(" \033[36m%d test(s) pending\033[0m\n\n", r.pending)
    }

    if len(r.failures) > 0 {
        fmt.Printf("%s \n\n", red(fmt.Sprintf(" %d tests failed:", len(r.failures))))

    }

    for i, failure := range r.failures {
        fmt.Printf("  %d) %s:\n\n", i+1, failure.testName)
        fmt.Printf("    %s %s\n\n", red(failure.message), gray(fmt.Sprintf("(%s:%d)", failure.file, failure.line)))
    }
}
