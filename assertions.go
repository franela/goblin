package goblin

import (
	"fmt"
	"reflect"
	"strings"
)

type Assertion struct {
	src  interface{}
	fail func(interface{})
}

func formatMessages(messages ...string) string {
	if len(messages) > 0 {
		return ", " + strings.Join(messages, " ")
	}
	return ""
}

func (a *Assertion) Eql(dst interface{}) {
	a.Equal(dst)
}

func (a *Assertion) Equal(dst interface{}) {
	if e := a.equal(dst); e != nil {
		a.fail(e)
	}
}

func (a *Assertion) equal(dst interface{}) error {
	if at, dt := reflect.TypeOf(a.src), reflect.TypeOf(dst); at != dt {
		return fmt.Errorf("%s(%#v) does not equal %s(%#v)", at, a.src, dt, dst)
	}

	if reflect.DeepEqual(a.src, dst) {
		return nil
	}

	if fmt.Sprintf("%#v", a.src) == fmt.Sprintf("%#v", dst) {
		return nil
	}

	return fmt.Errorf("%#v does not equal %#v", a.src, dst)
}

func (a *Assertion) IsTrue(messages ...string) {
	if a.equal(true) != nil {
		message := fmt.Sprintf("%v %s%s", a.src, "expected false to be truthy", formatMessages(messages...))
		a.fail(message)
	}
}

func (a *Assertion) IsFalse(messages ...string) {
	if a.equal(false) != nil {
		message := fmt.Sprintf("%v %s%s", a.src, "expected true to be falsey", formatMessages(messages...))
		a.fail(message)
	}
}
