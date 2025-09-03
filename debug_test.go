package utilities_test

import (
	"errors"
	utilities "github.com/dan-sherwin/go-utilities"
	"testing"
)

func TestLitterCheckErr(t *testing.T) {
	v := map[string]int{"a": 1}
	got := utilities.LitterCheckErr(v, nil)
	if got["a"] != 1 {
		t.Fatalf("LitterCheckErr with nil err: unexpected return %v", got)
	}
	v2 := []string{"x"}
	got2 := utilities.LitterCheckErr(v2, errors.New("boom"))
	if len(got2) != 1 || got2[0] != "x" {
		t.Fatalf("LitterCheckErr with err: unexpected return %v", got2)
	}
}
