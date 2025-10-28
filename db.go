package utilities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

type (
	// DbDSNConfig contains the fields used to construct a database DSN string.
	// Fields map to host, port, database name, user, password, SSL mode and timezone.
	DbDSNConfig struct {
		Server   string
		Port     int
		Name     string
		User     string
		Password string
		SSLMode  bool
		TimeZone string
	}
)

// DbDSN generates a database connection string (DSN) based on the provided configuration structure, including server, port, database name, user, password, SSL mode, and timezone. The SSL mode defaults to "enable" or "disable" based on the cfg.SSLMode flag.
func DbDSN(cfg DbDSNConfig) string {
	var sm string
	if cfg.SSLMode {
		sm = "enable"
	} else {
		sm = "disable"
	}
	connstr := fmt.Sprintf("host=%s dbname=%s sslmode=%s", cfg.Server, cfg.Name, sm)
	if cfg.Port != 0 {
		connstr = fmt.Sprintf("%s port=%d", connstr, cfg.Port)
	}
	if len(cfg.User) > 0 {
		connstr = fmt.Sprintf("%s user=%s", connstr, cfg.User)
	}
	if len(cfg.Password) > 0 {
		connstr = fmt.Sprintf("%s password=%s", connstr, cfg.Password)
	}
	if len(cfg.TimeZone) > 0 {
		connstr = fmt.Sprintf("%s TimeZone=%s", connstr, cfg.TimeZone)
	}
	return connstr
}

// ToValuers converts a slice of types implementing driver.Valuer into a slice of driver.Valuer.
// Useful for building slices for parameterized queries (e.g., WHERE IN (...)).
func ToValuers[T driver.Valuer](in []T) []driver.Valuer {
	out := make([]driver.Valuer, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

// ConvertToJSONMap marshals the input value to JSON and unmarshals it into a datatypes.JSONMap.
// Returns an error if marshaling or unmarshaling fails.
func ConvertToJSONMap(input any) (datatypes.JSONMap, error) {
	// Marshal the input to a JSON byte slice
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	jsonMap := datatypes.JSONMap{}
	// Unmarshal the JSON bytes into the JSONMap

	if err := jsonMap.UnmarshalJSON(bytes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal to JSONMap: %w", err)
	}

	return jsonMap, nil
}
