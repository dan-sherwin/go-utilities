package utilities

import (
	"database/sql/driver"
	"fmt"
)

type (
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

// ToValuers - Used in the creation of Where conditions
func ToValuers[T driver.Valuer](in []T) []driver.Valuer {
	out := make([]driver.Valuer, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}
