package goblin

import (
    "reflect"
    "fmt"
)

type Assertion struct {
  src interface{}
}

func objectsAreEqual(a, b interface{}) bool {
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

func (a *Assertion) Equal(dst interface{}) {
    if !objectsAreEqual(a.src, dst) {
        panic(fmt.Sprintf("%v", a.src)+" does not equal "+fmt.Sprintf("%v", dst))
    }
}
