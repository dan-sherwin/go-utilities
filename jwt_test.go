package utilities_test

import (
	utilities "github.com/dan-sherwin/go-utilities"
	"testing"
	"time"
)

type jwtClaims struct {
	UserID string `json:"uid"`
}

func TestJWTHelpers(t *testing.T) {
	secret := []byte("secret")
	claims := jwtClaims{UserID: "u1"}

	token, err := utilities.GenerateJWT(claims, time.Minute, secret)
	if err != nil {
		t.Fatalf("GenerateJWT error: %v", err)
	}
	if token == "" {
		t.Fatalf("GenerateJWT returned empty token")
	}

	if _, err := utilities.ValidateJWT(token, secret); err != nil {
		t.Fatalf("ValidateJWT error: %v", err)
	}

	m, err := utilities.ExtractJwtClaims(token, secret)
	if err != nil {
		t.Fatalf("ExtractJwtClaims error: %v", err)
	}
	if got := m["uid"]; got != "u1" {
		t.Fatalf("Extracted uid=%v want u1", got)
	}
	if _, ok := m["iat"]; !ok {
		t.Errorf("iat not present in claims")
	}
	if _, ok := m["exp"]; !ok {
		t.Errorf("exp not present in claims")
	}

	var out jwtClaims
	if err := utilities.ExtractJwtClaimsInto(token, secret, &out); err != nil {
		t.Fatalf("ExtractJwtClaimsInto error: %v", err)
	}
	if out.UserID != "u1" {
		t.Fatalf("ExtractJwtClaimsInto uid=%q want u1", out.UserID)
	}

	// wrong secret should fail
	wrong := []byte("wrong")
	if _, err := utilities.ValidateJWT(token, wrong); err == nil {
		t.Errorf("ValidateJWT with wrong secret should fail")
	}
	if _, err := utilities.ExtractJwtClaims(token, wrong); err == nil {
		t.Errorf("ExtractJwtClaims with wrong secret should fail")
	}
}
