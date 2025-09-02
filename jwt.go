package utilities

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateJWT creates a JSON Web Token (JWT) with the specified claims and expiration duration using the provided secret key.
// It embeds issued-at (iat) and expiration (exp) fields into the claims automatically.
// Returns the signed JWT string or an error if the operation fails.
func GenerateJWT(claims interface{}, duration time.Duration, secretKey []byte) (string, error) {
	raw, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	var claimsMap map[string]interface{}
	if err := json.Unmarshal(raw, &claimsMap); err != nil {
		return "", err
	}
	now := time.Now()
	claimsMap["iat"] = now.Unix()
	claimsMap["exp"] = now.Add(duration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claimsMap))
	return token.SignedString(secretKey)
}

// ValidateJWT validates a JWT using the provided token string and secret key, returning the parsed token if successful or an error if validation fails.
func ValidateJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return secretKey, nil
	}, jwt.WithValidMethods([]string{"HS256"}))
}

// ExtractJwtClaims extracts claims from a JWT token string using the provided secret key.
// It returns the claims as a map if the token is valid, or an error if validation or extraction fails.
func ExtractJwtClaims(tokenString string, secretKey []byte) (map[string]interface{}, error) {
	token, err := ValidateJWT(tokenString, secretKey)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}

// ExtractJwtClaimsInto extracts claims from a JWT token string, validates it using the provided secretKey, and maps the claims into the provided output struct. Returns an error if validation, marshalling, or unmarshalling fails.
func ExtractJwtClaimsInto(tokenString string, secretKey []byte, out interface{}) error {
	claimsMap, err := ExtractJwtClaims(tokenString, secretKey)
	if err != nil {
		return err
	}
	raw, err := json.Marshal(claimsMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}
