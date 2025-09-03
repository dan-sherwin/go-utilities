package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"testing"
)

func TestPtr(t *testing.T) {
	p := utilities.Ptr(42)
	if p == nil || *p != 42 {
		t.Fatalf("Ptr failed: got %v", p)
	}
}

func TestPtrZeroNil(t *testing.T) {
	if got := utilities.PtrZeroNil(0); got != nil {
		t.Errorf("PtrZeroNil(0) expected nil, got %v", got)
	}
	if got := utilities.PtrZeroNil(5); got == nil || *got != 5 {
		t.Errorf("PtrZeroNil(5) expected *5, got %v", got)
	}
	if got := utilities.PtrZeroNil(""); got != nil {
		t.Errorf("PtrZeroNil(\"\") expected nil, got %v", got)
	}
}

func TestPtrVal(t *testing.T) {
	if v := utilities.PtrVal[int](nil); v != 0 {
		t.Errorf("PtrVal(nil) expected 0, got %v", v)
	}
	p := utilities.Ptr("hi")
	if v := utilities.PtrVal(p); v != "hi" {
		t.Errorf("PtrVal(p) expected hi, got %v", v)
	}
}

func TestPtrCompare(t *testing.T) {
	if !utilities.PtrCompare[int](nil, nil) {
		t.Errorf("PtrCompare(nil,nil) expected true")
	}
	a, b := 1, 1
	if !utilities.PtrCompare(&a, &b) {
		t.Errorf("PtrCompare(&1,&1) expected true")
	}
	c := 2
	if utilities.PtrCompare(&a, &c) {
		t.Errorf("PtrCompare(&1,&2) expected false")
	}
	if utilities.PtrCompare(&a, nil) {
		t.Errorf("PtrCompare(&1,nil) expected false")
	}
}

func TestCopyIfNotNil(t *testing.T) {
	var dest int
	src := 7
	utilities.CopyIfNotNil(&src, &dest)
	if dest != 7 {
		t.Errorf("CopyIfNotNil failed, dest=%d want 7", dest)
	}
	// no panic when either is nil
	utilities.CopyIfNotNil[int](nil, &dest)
	if dest != 7 {
		t.Errorf("CopyIfNotNil with nil src should not modify dest")
	}
	utilities.CopyIfNotNil(&src, nil)
}

func TestCopyIfNotZero(t *testing.T) {
	var dest int
	utilities.CopyIfNotZero(0, &dest)
	if dest != 0 {
		t.Errorf("CopyIfNotZero with zero should not change dest, got %d", dest)
	}
	utilities.CopyIfNotZero(9, &dest)
	if dest != 9 {
		t.Errorf("CopyIfNotZero failed, dest=%d want 9", dest)
	}
}

func TestNilIfEmpty(t *testing.T) {
	if p := utilities.NilIfEmpty([]int{}); p != nil {
		t.Errorf("NilIfEmpty(empty) expected nil, got %v", p)
	}
	s := []string{"a"}
	p := utilities.NilIfEmpty(s)
	if p == nil || len(*p) != 1 || (*p)[0] != "a" {
		t.Errorf("NilIfEmpty(non-empty) unexpected result: %#v", p)
	}
}

func TestNilIfZeroPtr(t *testing.T) {
	var x int
	px := &x
	if got := utilities.NilIfZeroPtr(px); got != nil {
		t.Errorf("NilIfZeroPtr(&0) expected nil, got %v", got)
	}
	y := 3
	py := &y
	if got := utilities.NilIfZeroPtr(py); got != py {
		t.Errorf("NilIfZeroPtr(&3) expected same pointer, got %v", got)
	}
	if got := utilities.NilIfZeroPtr[int](nil); got != nil {
		t.Errorf("NilIfZeroPtr(nil) expected nil, got %v", got)
	}
}
