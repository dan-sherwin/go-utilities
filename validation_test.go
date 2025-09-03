package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"strings"
	"testing"
)

func TestIsEmail(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"alice@example.com", true},
		{" Bob <bob@example.org> ", true},
		{"no-at-symbol", false},
		{"foo@", false},
		{"@bar", false},
		{"user@localhost", false}, // not FQDN (no dot), allowed by hostname but IsEmail requires FQDN or IP
		{"user@127.0.0.1", true},
	}
	for _, c := range cases {
		if got := utilities.IsEmail(c.in); got != c.ok {
			t.Errorf("IsEmail(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsIP(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"127.0.0.1", true},
		{"::1", true},
		{"256.0.0.1", false},
		{"gibberish", false},
	}
	for _, c := range cases {
		if got := utilities.IsIP(c.in); got != c.ok {
			t.Errorf("IsIP(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsIPv4v6(t *testing.T) {
	if !utilities.IsIPv4("8.8.8.8") {
		t.Errorf("IsIPv4 failed for 8.8.8.8")
	}
	if utilities.IsIPv4("::1") {
		t.Errorf("IsIPv4 should be false for ::1")
	}
	if !utilities.IsIPv6("2001:db8::1") {
		t.Errorf("IsIPv6 failed for 2001:db8::1")
	}
	if utilities.IsIPv6("192.168.0.1") {
		t.Errorf("IsIPv6 should be false for IPv4")
	}
}

func TestIsMAC(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"00:1a:2b:3c:4d:5e", true},
		{"00-1A-2B-3C-4D-5E", true},
		{"00:1a:2b:3c:4d", false}, // too short (5 groups)
		{"zz:zz:zz:zz:zz:zz", false},
	}
	for _, c := range cases {
		if got := utilities.IsMAC(c.in); got != c.ok {
			t.Errorf("IsMAC(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsURL(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"https://example.com", true},
		{"http://example.com/path?x=y", true},
		{"https://[2001:db8::1]", true},
		{"ftp://example.com", false},
		{"http:///nohost", false},
		{"http://", false},
	}
	for _, c := range cases {
		if got := utilities.IsURL(c.in); got != c.ok {
			t.Errorf("IsURL(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsFQDN(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"example.com", true},
		{"sub.example.com", true},
		{"-bad.example", false},
		{"example-.com", false},
		{"localhost", false},
		{"a..b", false},
	}
	for _, c := range cases {
		if got := utilities.IsFQDN(c.in); got != c.ok {
			t.Errorf("IsFQDN(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsHostname(t *testing.T) {
	long := strings.Repeat("a", 64)
	cases := []struct {
		in string
		ok bool
	}{
		{"localhost", true},
		{"my-host", true},
		{"-bad", false},
		{"bad-", false},
		{long, false},
	}
	for _, c := range cases {
		if got := utilities.IsHostname(c.in); got != c.ok {
			t.Errorf("IsHostname(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsCreditCard(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"4242 4242 4242 4242", true},
		{"4012-8888-8888-1881", true},
		{"1234 5678 9012 3456", false},
		{"abcd efgh ijkl mnop", false},
		{"  ", false},
	}
	for _, c := range cases {
		if got := utilities.IsCreditCard(c.in); got != c.ok {
			t.Errorf("IsCreditCard(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsPhone(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{in: "+15551234567", ok: true},
		{in: "5551234567", ok: true},
		{in: "+0123456", ok: false}, // leading 0 after + not allowed
		{in: "123", ok: false},
	}
	for _, c := range cases {
		if got := utilities.IsPhone(c.in); got != c.ok {
			t.Errorf("IsPhone(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsUUID(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"550e8400-e29b-41d4-a716-446655440000", true}, // v4
		{"550e8400e29b41d4a716446655440000", false},
		{"not-a-uuid", false},
	}
	for _, c := range cases {
		if got := utilities.IsUUID(c.in); got != c.ok {
			t.Errorf("IsUUID(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsBase64(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"SGVsbG8gd29ybGQ=", true},
		{"SGVsbG8gd29ybGQ", true}, // raw no padding
		{"!! not base64 !!", false},
	}
	for _, c := range cases {
		if got := utilities.IsBase64(c.in); got != c.ok {
			t.Errorf("IsBase64(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsHexColor(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"#fff", true},
		{"#ffffff", true},
		{"#abcd", true},
		{"#aabbccdd", true},
		{"fff", true},
		{"#ggg", false},
	}
	for _, c := range cases {
		if got := utilities.IsHexColor(c.in); got != c.ok {
			t.Errorf("IsHexColor(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	cases := []struct {
		in string
		ok bool
	}{
		{"abc123", true},
		{"ABCxyz789", true},
		{"with space", false},
		{"punctuation!", false},
	}
	for _, c := range cases {
		if got := utilities.IsAlphaNumeric(c.in); got != c.ok {
			t.Errorf("IsAlphaNumeric(%q)=%v want %v", c.in, got, c.ok)
		}
	}
}
