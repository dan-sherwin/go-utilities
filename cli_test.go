package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"io"
	"os"
	"testing"
)

type user struct {
	ID   int
	Name string
}

func captureStdout(fn func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := fn()
	if cerr := w.Close(); cerr != nil && err == nil {
		// prefer original err if present
		err = cerr
	}
	os.Stdout = old
	b, rerr := io.ReadAll(r)
	if rerr != nil && err == nil {
		err = rerr
	}
	return string(b), err
}

func TestCLI_PrintStructTable(t *testing.T) {
	out, err := captureStdout(func() error {
		return utilities.PrintStructTable([]user{{1, "Ada"}, {2, "Linus"}})
	})
	if err != nil {
		t.Fatalf("PrintStructTable error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintStructTable")
	}
}

func TestCLI_PrintMapArray(t *testing.T) {
	rows := []map[string]any{{"id": 1, "name": "Ada"}, {"id": 2, "name": "Linus"}}
	out, err := captureStdout(func() error { return utilities.PrintMapArray(rows) })
	if err != nil {
		t.Fatalf("PrintMapArray error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintMapArray")
	}
}

func TestCLI_PrintStructMap_and_Sorted(t *testing.T) {
	m := map[string]user{"a": {1, "Ada"}, "b": {2, "Linus"}}
	out, err := captureStdout(func() error { return utilities.PrintStructMap(m) })
	if err != nil {
		t.Fatalf("PrintStructMap error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintStructMap")
	}
	out, err = captureStdout(func() error { return utilities.PrintSortedStructMap(m) })
	if err != nil {
		t.Fatalf("PrintSortedStructMap error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintSortedStructMap")
	}
}

func TestCLI_PrintStringSlice(t *testing.T) {
	out, err := captureStdout(func() error { return utilities.PrintStringSlice([]string{"a", "b"}) })
	if err != nil {
		t.Fatalf("PrintStringSlice error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintStringSlice")
	}
}

func TestCLI_PrintAnySlice(t *testing.T) {
	out, err := captureStdout(func() error { return utilities.PrintAnySlice([]any{1, "two", 3.0}) })
	if err != nil {
		t.Fatalf("PrintAnySlice error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintAnySlice")
	}
}

func TestCLI_PrintSlice(t *testing.T) {
	out, err := captureStdout(func() error { return utilities.PrintSlice([]int{1, 2, 3}) })
	if err != nil {
		t.Fatalf("PrintSlice error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintSlice")
	}
}

func TestCLI_PrintMap(t *testing.T) {
	m := map[string]any{"id": 1, "name": "Ada"}
	out, err := captureStdout(func() error { return utilities.PrintMap(m, "id", "name") })
	if err != nil {
		t.Fatalf("PrintMap error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintMap")
	}
}

func TestCLI_PrintStringsTable(t *testing.T) {
	headers := []string{"ID", "Name"}
	rows := [][]string{{"1", "Ada"}, {"2", "Linus"}}
	out, err := captureStdout(func() error { return utilities.PrintStringsTable(headers, rows) })
	if err != nil {
		t.Fatalf("PrintStringsTable error: %v", err)
	}
	if len(out) == 0 {
		t.Errorf("no output from PrintStringsTable")
	}
}
