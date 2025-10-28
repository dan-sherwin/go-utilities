package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"os"
	"path/filepath"
	"testing"
)

func TestDirCreateIfNotExists(t *testing.T) {
	tmp := t.TempDir()
	dir := filepath.Join(tmp, "nested", "path")
	// should create without error
	if err := utilities.DirCreateIfNotExists(dir); err != nil {
		t.Fatalf("DirCreateIfNotExists error: %v", err)
	}
	st, err := os.Stat(dir)
	if err != nil || !st.IsDir() {
		t.Fatalf("expected directory to exist: err=%v, stat=%v", err, st)
	}
	// calling again should be idempotent
	if err := utilities.DirCreateIfNotExists(dir); err != nil {
		t.Fatalf("DirCreateIfNotExists second call error: %v", err)
	}
}

func TestAmAdmin_Smoke(t *testing.T) {
	t.Helper()
	_ = utilities.AmAdmin() // just ensure it doesn't panic and returns a bool
}
