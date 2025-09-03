package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	a := map[string]int{"a": 1}
	b := map[string]int{"a": 2, "b": 3}
	out := utilities.Merge(a, b)
	// overwrite from b
	if out["a"] != 2 || out["b"] != 3 {
		t.Errorf("unexpected merged map: %#v", out)
	}
	// originals unchanged
	if a["a"] != 1 {
		t.Errorf("original map a modified: %#v", a)
	}
	if _, ok := a["b"]; ok {
		t.Errorf("original map a gained key 'b': %#v", a)
	}

	// nil inputs
	var nilMap map[string]int
	out2 := utilities.Merge(nilMap, b)
	if !reflect.DeepEqual(out2, b) {
		t.Errorf("Merge(nil,b) => %v want %v", out2, b)
	}
	out3 := utilities.Merge(a, nilMap)
	if !reflect.DeepEqual(out3, a) {
		t.Errorf("Merge(a,nil) => %v want %v", out3, a)
	}
}

func TestMergeInto(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 3, "z": 4}
	utilities.MergeInto(a, b)
	if a["x"] != 1 || a["y"] != 3 || a["z"] != 4 {
		t.Errorf("MergeInto result unexpected: %#v", a)
	}
	// b unchanged
	if !reflect.DeepEqual(b, map[string]int{"y": 3, "z": 4}) {
		t.Errorf("source map b should remain unchanged: %#v", b)
	}
	// merging nil b is a no-op
	utilities.MergeInto[string, int](a, nil)
	if a["x"] != 1 || a["y"] != 3 || a["z"] != 4 {
		t.Errorf("MergeInto with nil b should not change a: %#v", a)
	}
}
