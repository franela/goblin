package goblin

import (
	"fmt"
	"reflect"
)

type Assertion interface {
	Eql(dst interface{})
	Equal(dst interface{})
	IsTrue()
	IsFalse()
	SetSource(interface{})
}

type AssertionV1 struct {
	src  interface{}
	fail func(interface{})
}

type AssertionV2 struct {
	src  interface{}
	fail func(interface{})
}

func objectsAreEqual(a, b interface{}) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	if reflect.DeepEqual(a, b) {
		return true
	}

	if reflect.ValueOf(a) == reflect.ValueOf(b) {
		return true
	}

	if fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b) {
		return true
	}

	return false
}

func (a *AssertionV1) SetSource(src interface{}) {
	a.src = src
}

func (a *AssertionV1) Eql(dst interface{}) {
	a.Equal(dst)
}

func (a *AssertionV1) Equal(dst interface{}) {
	if !objectsAreEqual(a.src, dst) {
		a.fail(fmt.Sprintf("%v", a.src) + " does not equal " + fmt.Sprintf("%v", dst))
	}
}

func (a *AssertionV1) IsTrue() {
	if !objectsAreEqual(a.src, true) {
		a.fail(fmt.Sprintf("%v", a.src) + " expected false to be truthy")
	}
}

func (a *AssertionV1) IsFalse() {
	if !objectsAreEqual(a.src, false) {
		a.fail(fmt.Sprintf("%v", a.src) + " expected true to be falsey")
	}
}

func (a *AssertionV2) SetFailFunc(fail func(interface{})) {
	a.fail = fail
}

func (a *AssertionV2) SetSource(src interface{}) {
	a.src = src
}

func (a *AssertionV2) Eql(dst interface{}) {
	a.Equal(dst)
}

func (a *AssertionV2) Equal(dst interface{}) {
	if !objectsAreEqual(a.src, dst) {
		a.fail(fmt.Sprintf("%v", a.src) + " does not equal " + fmt.Sprintf("%v", dst))
	}
}

func (a *AssertionV2) IsTrue() {
	a.fail("Not implemented in assertion v2, use Assert(bool) instead")
}

func (a *AssertionV2) IsFalse() {
	a.fail("Not implemented in assertion v2, use Assert(bool) instead")
}
