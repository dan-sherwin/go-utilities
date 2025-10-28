package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"testing"
)

type sample struct {
	A int
	B string
	C *int
}

func TestZeroStructFieldByName(t *testing.T) {
	s := sample{A: 10, B: "x"}
	if err := utilities.ZeroStructFieldByName(&s, "A"); err != nil {
		t.Fatalf("ZeroStructFieldByName error: %v", err)
	}
	if s.A != 0 {
		t.Errorf("expected A zeroed, got %d", s.A)
	}
	// error on non-existing field
	if err := utilities.ZeroStructFieldByName(&s, "Z"); err == nil {
		t.Errorf("expected error on missing field, got nil")
	}
	// error when not passing pointer-to-struct
	if err := utilities.ZeroStructFieldByName(s, "A"); err == nil {
		t.Errorf("expected error when not pointer to struct")
	}
}

func TestSetStructFieldByName(t *testing.T) {
	s := sample{}
	if err := utilities.SetStructFieldByName(&s, "A", 5); err != nil {
		t.Fatalf("SetStructFieldByName error: %v", err)
	}
	if s.A != 5 {
		t.Errorf("expected A=5, got %d", s.A)
	}
	if err := utilities.SetStructFieldByName(&s, "B", "hello"); err != nil {
		t.Fatalf("SetStructFieldByName error: %v", err)
	}
	if s.B != "hello" {
		t.Errorf("expected B=hello, got %q", s.B)
	}
	// type mismatch
	if err := utilities.SetStructFieldByName(&s, "A", "not-int"); err == nil {
		t.Errorf("expected type mismatch error, got nil")
	}
	// missing field
	if err := utilities.SetStructFieldByName(&s, "Z", 1); err == nil {
		t.Errorf("expected missing field error, got nil")
	}
	// not pointer to struct
	if err := utilities.SetStructFieldByName(s, "A", 1); err == nil {
		t.Errorf("expected error when not pointer to struct")
	}
}

func TestStructFieldNames(t *testing.T) {
	s := sample{}
	names := utilities.StructFieldNames(s)
	if len(names) != 3 {
		t.Fatalf("expected 3 names, got %v", names)
	}
	// ensure contains A, B, C
	found := map[string]bool{}
	for _, n := range names {
		found[n] = true
	}
	if !found["A"] || !found["B"] || !found["C"] {
		t.Errorf("missing expected fields in %v", names)
	}
}

func TestStructToStringMap(t *testing.T) {
	v := 42
	s := sample{A: 1, B: "ok", C: &v}
	m := utilities.StructToStringMap(s)
	if m["A"] != "1" || m["B"] != "ok" || m["C"] != "42" {
		t.Errorf("unexpected map values: %#v", m)
	}
	// nil pointer shows as <nil>
	s2 := sample{}
	m2 := utilities.StructToStringMap(&s2)
	if m2["C"] != "<nil>" {
		t.Errorf("expected C to be <nil>, got %q", m2["C"])
	}
}
