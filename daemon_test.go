package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"testing"
)

func TestDaemonHelpers_NegativeCases(t *testing.T) {
	name := "this-process-name-should-not-exist-xyz"
	if utilities.DaemonAlreadyRunning(name) {
		// Extremely unlikely, but avoid flakiness
		t.Skip("unexpectedly found a process with the test name; skipping")
	}
	if pid := utilities.MustFindDaemonProcessPID(name); pid != 0 {
		t.Errorf("MustFindDaemonProcessPID expected 0 for not found, got %d", pid)
	}
	if _, err := utilities.FindDaemonProcessPID(name); err == nil {
		t.Errorf("FindDaemonProcessPID expected error for not found process")
	}
}
