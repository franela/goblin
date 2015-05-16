package goblin

import (
	"fmt"
	"reflect"
	"strings"
)

type Assertion struct {
	src  interface{}
	fail func(interface{})
	Not  *NotAssertion
}

func objectsAreEqual(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	if reflect.DeepEqual(a, b) {
		return true
	}

	if fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b) {
		return true
	}

	return false
}

func formatMessages(messages ...string) string {
	if len(messages) > 0 {
		return ", " + strings.Join(messages, " ")
	}
	return ""
}

func (a *Assertion) Eql(dst interface{}, messages ...string) {
	a.Equal(dst, messages...)
}

func (a *Assertion) Equal(dst interface{}, messages ...string) {
	if !objectsAreEqual(a.src, dst) {
		a.fail(fmt.Sprintf("%v does not equal %v%s", a.src, dst, formatMessages(messages...)))
	}
}

func (a *Assertion) IsTrue(messages ...string) {
	if !objectsAreEqual(a.src, true) {
		a.fail(fmt.Sprintf("Expected %v to be truthy%s", a.src, formatMessages(messages...)))
	}
}

func (a *Assertion) IsFalse(messages ...string) {
	if !objectsAreEqual(a.src, false) {
		a.fail(fmt.Sprintf("Expected %v to be falsey%s", a.src, formatMessages(messages...)))
	}
}

func (a *Assertion) IsOK(messages ...string) {
	if objectsAreEqual(a.src, nil) {
		a.fail(fmt.Sprintf("Expected %v to be OK%s", a.src, formatMessages(messages...)))
	}
}

func (a *Assertion) HasSameType(dst interface{}, messages ...string) {
	if reflect.TypeOf(a.src) != reflect.TypeOf(dst) {
		a.fail(fmt.Sprintf("%v is not the same type as %v%s", a.src, dst, formatMessages(messages...)))
	}

}

type NotAssertion struct {
	src  interface{}
	fail func(interface{})
	Not  *Assertion
}

func (n *NotAssertion) Eql(dst interface{}, messages ...string) {
	n.Equal(dst, messages...)
}

func (n *NotAssertion) Equal(dst interface{}, messages ...string) {
	if objectsAreEqual(n.src, dst) {
		n.fail(fmt.Sprintf("Expected %v to not equal %v%s", n.src, dst, formatMessages(messages...)))
	}
}

func (n *NotAssertion) IsTrue(messages ...string) {
	if objectsAreEqual(n.src, true) {
		n.fail(fmt.Sprintf("Expected %v to not be truthy%s", n.src, formatMessages(messages...)))
	}
}

func (n *NotAssertion) IsFalse(messages ...string) {
	if objectsAreEqual(n.src, false) {
		n.fail(fmt.Sprintf("Expected %v to not be falsey%s", n.src, formatMessages(messages...)))
	}
}

func (n *NotAssertion) IsOK(messages ...string) {
	if !objectsAreEqual(n.src, nil) {
		n.fail(fmt.Sprintf("Expected %v to not be OK%s", n.src, formatMessages(messages...)))
	}
}

func (n *NotAssertion) HasSameType(dst interface{}, messages ...string) {
	if reflect.TypeOf(n.src) == reflect.TypeOf(dst) {
		n.fail(fmt.Sprintf("Expected %v to not be the same type as %v%s", n.src, dst, formatMessages(messages...)))
	}
}
