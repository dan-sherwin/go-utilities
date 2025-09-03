package utilities_test

import (
	"database/sql/driver"
	utilities "github.com/dan-sherwin/go-utilities"
	"testing"
)

type myValuer int

func (m myValuer) Value() (driver.Value, error) { return int(m), nil }

func TestDbDSN(t *testing.T) {
	cfg := utilities.DbDSNConfig{Server: "localhost", Name: "testdb", SSLMode: true}
	if got, want := utilities.DbDSN(cfg), "host=localhost dbname=testdb sslmode=enable"; got != want {
		t.Fatalf("DbDSN basic mismatch: got %q want %q", got, want)
	}

	cfg = utilities.DbDSNConfig{Server: "db", Port: 5433, Name: "app", User: "u", Password: "p", SSLMode: false, TimeZone: "UTC"}
	want := "host=db dbname=app sslmode=disable port=5433 user=u password=p TimeZone=UTC"
	if got := utilities.DbDSN(cfg); got != want {
		t.Fatalf("DbDSN full mismatch: got %q want %q", got, want)
	}
}

func TestToValuers(t *testing.T) {
	in := []myValuer{1, 2, 3}
	vals := utilities.ToValuers(in)
	if len(vals) != 3 {
		t.Fatalf("ToValuers length=%d want 3", len(vals))
	}
	for i, v := range vals {
		got, err := v.Value()
		if err != nil {
			t.Fatalf("valuer %d returned error: %v", i, err)
		}
		switch n := got.(type) {
		case int64:
			if int(n) != int(in[i]) {
				t.Fatalf("unexpected int64 value at %d: %#v", i, got)
			}
		case int:
			if n != int(in[i]) {
				t.Fatalf("unexpected int value at %d: %#v", i, got)
			}
		default:
			t.Fatalf("unexpected driver.Value type %T at %d: %#v", got, i, got)
		}
	}
}
