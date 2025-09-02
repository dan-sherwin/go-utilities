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
		return nil, fmt.Errorf("Invalid credentials")
	}
	bearerToken := bearerHeader[7:]
	return utilities.ExtractJwtClaims(bearerToken, secretKey)
}

// ExtractJwtClaimsFromContextInto extracts JWT claims from the Authorization header of the given gin.Context, validates the token using the provided secretKey, and decodes the claims into the out parameter. An error is returned if the process fails at any step.
func ExtractJwtClaimsFromContextInto(c *gin.Context, secretKey []byte, out interface{}) error {
	bearerHeader := c.GetHeader("Authorization")
	if len(bearerHeader) < 8 {
		return fmt.Errorf("Invalid credentials")
	}
	bearerToken := bearerHeader[7:]
	return utilities.ExtractJwtClaimsInto(bearerToken, secretKey, out)
}
