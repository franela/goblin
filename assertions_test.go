package goblin

import (
	"testing"
)

// So we can test asserting type against its type alias
type String string

// Helper for testing Assertion conditions
type AssertionVerifier struct {
	ShouldPass bool
	didFail    bool
}

func (a *AssertionVerifier) FailFunc(msg interface{}) {
	a.didFail = true
}

func (a *AssertionVerifier) Verify(t *testing.T) {
	if a.didFail == a.ShouldPass {
		t.FailNow()
	}
}

func TestEqual(t *testing.T) {

	verifier := AssertionVerifier{ShouldPass: true}
	a := AssertionV1{src: 1, fail: verifier.FailFunc}
	a.Equal(1)
	verifier.Verify(t)
	a.Eql(1)
	verifier.Verify(t)

	a = AssertionV1{src: "foo"}
	a.Equal("foo")
	verifier.Verify(t)
	a.Eql("foo")
	verifier.Verify(t)

	a = AssertionV1{src: map[string]string{"foo": "bar"}}
	a.Equal(map[string]string{"foo": "bar"})
	verifier.Verify(t)
	a.Eql(map[string]string{"foo": "bar"})
	verifier.Verify(t)

	verifier = AssertionVerifier{ShouldPass: false}
	a = AssertionV1{src: String("baz"), fail: verifier.FailFunc}
	a.Equal("baz")
	verifier.Verify(t)
	a.Eql("baz")
	verifier.Verify(t)
}

func TestIsTrue(t *testing.T) {
	verifier := AssertionVerifier{ShouldPass: true}
	a := AssertionV1{src: true, fail: verifier.FailFunc}
	a.IsTrue()
	verifier.Verify(t)

	verifier = AssertionVerifier{ShouldPass: false}
	a = AssertionV1{src: false, fail: verifier.FailFunc}
	a.IsTrue()
	verifier.Verify(t)
}

func TestIsFalse(t *testing.T) {
	verifier := AssertionVerifier{ShouldPass: true}
	a := AssertionV1{src: false, fail: verifier.FailFunc}
	a.IsFalse()
	verifier.Verify(t)

	verifier = AssertionVerifier{ShouldPass: false}
	a = AssertionV1{src: true, fail: verifier.FailFunc}
	a.IsFalse()
	verifier.Verify(t)
}
