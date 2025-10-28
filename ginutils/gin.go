// Package ginutils provides helpers for extracting and decoding JWTs from gin.Context.
package ginutils

import (
	"fmt"

	"github.com/dan-sherwin/go-utilities"
	"github.com/gin-gonic/gin"
)

// ExtractJwtClaimsFromContext extracts JWT claims from the request context by parsing the Authorization header and decoding the token using the provided secret key. Returns a map of claims or an error if token parsing fails.
func ExtractJwtClaimsFromContext(c *gin.Context, secretKey []byte) (map[string]interface{}, error) {
	bearerHeader := c.GetHeader("Authorization")
	if len(bearerHeader) < 8 {
		return nil, fmt.Errorf("invalid credentials")
	}
	bearerToken := bearerHeader[7:]
	return utilities.ExtractJwtClaims(bearerToken, secretKey)
}

// ExtractJwtClaimsFromContextInto extracts JWT claims from the Authorization header of the given gin.Context, validates the token using the provided secretKey, and decodes the claims into the out parameter. An error is returned if the process fails at any step.
func ExtractJwtClaimsFromContextInto(c *gin.Context, secretKey []byte, out interface{}) error {
	bearerHeader := c.GetHeader("Authorization")
	if len(bearerHeader) < 8 {
		return fmt.Errorf("invalid credentials")
	}
	bearerToken := bearerHeader[7:]
	return utilities.ExtractJwtClaimsInto(bearerToken, secretKey, out)
}

// ExtractCookieJwtClaimsFromContext retrieves JWT claims from a cookie named "cxjwt" in the provided Gin context.
// The JWT is validated and parsed using the provided secret key.
// It returns the claims as a map if successful, or an error if the process fails.
func ExtractCookieJwtClaimsFromContext(c *gin.Context, secretKey []byte) (map[string]interface{}, error) {
	token, err := c.Cookie("cxjwt")
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	return utilities.ExtractJwtClaims(token, secretKey)
}

// ExtractCookieJwtClaimsFromContextInto retrieves a JWT from the "cxjwt" cookie in the provided gin.Context, validates it using the secretKey, and unmarshals the claims into the provided output struct. Returns an error if the cookie is not present or if validation/unmarshalling fails.
func ExtractCookieJwtClaimsFromContextInto(c *gin.Context, secretKey []byte, out interface{}) error {
	token, err := c.Cookie("cxjwt")
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}
	return utilities.ExtractJwtClaimsInto(token, secretKey, out)
}
