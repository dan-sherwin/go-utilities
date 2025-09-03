package utilities

import (
	"encoding/base64"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

// IsEmail returns true if s is a valid email address per net/mail parsing.
// It accepts common real-world emails but does not guarantee full RFC compliance.
func IsEmail(s string) bool {
	if s = strings.TrimSpace(s); s == "" {
		return false
	}
	addr, err := mail.ParseAddress(s)
	if err != nil {
		return false
	}
	// Ensure address has a local@domain form and the domain is plausible
	parts := strings.Split(addr.Address, "@")
	if len(parts) != 2 {
		return false
	}
	if parts[0] == "" || parts[1] == "" {
		return false
	}
	return IsFQDN(parts[1]) || IsIP(parts[1])
}

// IsIP returns true if s is a valid IPv4 or IPv6 address literal.
func IsIP(s string) bool { return net.ParseIP(strings.TrimSpace(s)) != nil }

// IsIPv4 returns true if s is a valid IPv4 address literal.
func IsIPv4(s string) bool {
	ip := net.ParseIP(strings.TrimSpace(s))
	return ip != nil && ip.To4() != nil
}

// IsIPv6 returns true if s is a valid IPv6 address literal.
func IsIPv6(s string) bool {
	ip := net.ParseIP(strings.TrimSpace(s))
	return ip != nil && ip.To4() == nil
}

// IsMAC returns true if s is a valid MAC-48 or EUI-64 hardware address.
func IsMAC(s string) bool { _, err := net.ParseMAC(strings.TrimSpace(s)); return err == nil }

// IsURL returns true if s is a valid HTTP or HTTPS URL with a host.
func IsURL(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	u, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	if u.Host == "" {
		return false
	}
	// Validate host part: allow IPv4, IPv6 in brackets, or FQDN/hostname
	host := u.Host
	// Strip port if present
	if i := strings.LastIndex(host, ":"); i != -1 {
		// IPv6 hosts will be in [::1]:443, keep brackets part for detection
		if !strings.Contains(host, "]") || i > strings.LastIndex(host, "]") {
			host = host[:i]
		}
	}
	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		return IsIPv6(strings.Trim(host, "[]"))
	}
	return IsIPv4(host) || IsFQDN(host) || IsHostname(host)
}

// IsFQDN validates a fully-qualified domain name per common DNS label rules.
// - total length <= 253
// - labels 1..63 chars, alnum and hyphens, not starting/ending with hyphen
// - at least one dot
var fqdnLabel = regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`)

func IsFQDN(s string) bool {
	s = strings.TrimSuffix(strings.TrimSpace(s), ".") // tolerate trailing dot but not required
	if s == "" || len(s) > 253 || !strings.Contains(s, ".") {
		return false
	}
	parts := strings.Split(s, ".")
	for _, p := range parts {
		if len(p) == 0 || len(p) > 63 || !fqdnLabel.MatchString(p) {
			return false
		}
	}
	return true
}

// IsHostname allows single-label hostnames (non-FQDN), same label rules as FQDN labels.
func IsHostname(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" || len(s) > 63 {
		return false
	}
	return fqdnLabel.MatchString(s)
}

// IsCreditCard performs a format check (digits with optional separators) and Luhn checksum.
func IsCreditCard(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	// Remove common separators
	normalized := make([]rune, 0, len(s))
	for _, r := range s {
		if r >= '0' && r <= '9' {
			normalized = append(normalized, r)
		} else if r == ' ' || r == '-' {
			continue
		} else {
			return false
		}
	}
	if l := len(normalized); l < 12 || l > 19 { // common card number lengths
		return false
	}
	// Luhn check
	sum := 0
	double := false
	for i := len(normalized) - 1; i >= 0; i-- {
		d := int(normalized[i] - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}

// IsPhone returns true for basic E.164 numbers: + followed by 8..15 digits, or national digits 7..15.
// This is intentionally conservative and does not validate per-country rules.
var (
	rePhoneE164  = regexp.MustCompile(`^\+[1-9][0-9]{7,14}$`)
	rePhoneLocal = regexp.MustCompile(`^[0-9]{7,15}$`)
)

func IsPhone(s string) bool {
	s = strings.TrimSpace(s)
	return rePhoneE164.MatchString(s) || rePhoneLocal.MatchString(s)
}

// IsUUID supports canonical UUID v1-5 (8-4-4-4-12 hex with hyphens).
var reUUID = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[aAbB89][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

func IsUUID(s string) bool { return reUUID.MatchString(strings.TrimSpace(s)) }

// IsBase64 checks if s is valid base64 without requiring specific padding, ignoring whitespace.
func IsBase64(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	// Remove common whitespace
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, " ", "")
	_, err := base64.StdEncoding.DecodeString(s)
	if err == nil {
		return true
	}
	_, err = base64.RawStdEncoding.DecodeString(s)
	return err == nil
}

// IsHexColor validates #RGB, #RRGGBB, #RGBA, #RRGGBBAA or without leading #.
var reHexColor = regexp.MustCompile(`^(#?)(?:[A-Fa-f0-9]{3}|[A-Fa-f0-9]{6}|[A-Fa-f0-9]{4}|[A-Fa-f0-9]{8})$`)

func IsHexColor(s string) bool { return reHexColor.MatchString(strings.TrimSpace(s)) }

// IsAlphaNumeric returns true if s contains only ASCII letters and digits.
var reAlphaNum = regexp.MustCompile(`^[A-Za-z0-9]+$`)

func IsAlphaNumeric(s string) bool { return reAlphaNum.MatchString(strings.TrimSpace(s)) }
