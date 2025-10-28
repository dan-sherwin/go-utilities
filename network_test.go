package utilities_test

import (
	"strings"
	"testing"

	utilities "github.com/dan-sherwin/go-utilities"
)

func TestGetMacAddressFromIp_InvalidAndNotFound(t *testing.T) {
	if _, err := utilities.GetMacAddressFromIp("not-an-ip"); err == nil {
		t.Fatalf("expected error for invalid IP input")
	} else if !strings.Contains(err.Error(), "invalid TV IP address") {
		t.Errorf("unexpected error message: %v", err)
	}
	// Use TEST-NET-1 address which should not be in local ARP cache in typical CI
	if _, err := utilities.GetMacAddressFromIp("192.0.2.1"); err == nil {
		t.Fatalf("expected error for non-present ARP entry")
	}
}
