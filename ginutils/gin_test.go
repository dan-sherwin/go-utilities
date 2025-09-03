package ginutils_test

import (
	"net/http/httptest"
	"testing"
	"time"

	utilities "github.com/dan-sherwin/go-utilities"
	"github.com/dan-sherwin/go-utilities/ginutils"
	"github.com/gin-gonic/gin"
)

type jwtClaims struct {
	UserID string `json:"uid"`
}

func TestExtractJwtFromContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	secret := []byte("secret")
	token, err := utilities.GenerateJWT(jwtClaims{UserID: "u1"}, time.Second, secret)
	if err != nil {
		t.Fatalf("GenerateJWT error: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	claims, err := ginutils.ExtractJwtClaimsFromContext(c, secret)
	if err != nil {
		t.Fatalf("ExtractJwtClaimsFromContext error: %v", err)
	}
	if claims["uid"] != "u1" {
		t.Fatalf("expected uid=u1, got %v", claims["uid"])
	}

	var out jwtClaims
	if err := ginutils.ExtractJwtClaimsFromContextInto(c, secret, &out); err != nil {
		t.Fatalf("ExtractJwtClaimsFromContextInto error: %v", err)
	}
	if out.UserID != "u1" {
		t.Fatalf("expected out.UserID=u1, got %q", out.UserID)
	}
}

func TestExtractJwtFromContext_InvalidHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	if _, err := ginutils.ExtractJwtClaimsFromContext(c, []byte("secret")); err == nil {
		t.Fatalf("expected error for missing/short Authorization header")
	}
}
