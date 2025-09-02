# Utilities for Go

A small, focused collection of helper functions and type shortcuts frequently used across services. Utilities are grouped around a few common themes:

- Pointers and generics helpers
- Struct reflection helpers (read/write fields, mapping, tabular output)
- CLI/table printing helpers
- JWT helpers (create, validate, extract, gin integration)
- Process/daemon helpers (PID discovery)
- Host/filesystem helpers
- File utilities (MIME type)
- Network helpers (ARP lookup)
- Database DSN helper and driver.Valuer utilities
- Common type aliases

This repository contains:
- Root package: github.com/dan-sherwin/go-utilities
- Subpackage: github.com/dan-sherwin/go-utilities/ginutils (gin-only JWT conveniences)

## Installation

- Root package:
  go get github.com/dan-sherwin/go-utilities

- Gin helpers:
  go get github.com/dan-sherwin/go-utilities/ginutils

Go version: see go.mod (currently go 1.23).

## Quick Start

- Create a JWT:
  
  token, err := utilities.GenerateJWT(struct{UserID string}{"123"}, time.Hour, []byte("secret"))
  
- Validate & read claims:
  
  claims, err := utilities.ExtractJwtClaims(token, []byte("secret"))
  
- Pointers:
  
  sPtr := utilities.Ptr("hello")        // *string
  v := utilities.PtrVal(sPtr)            // "hello"
  nonNil := utilities.PtrZeroNil(5)      // *int
  isEq := utilities.PtrCompare(&v, &v)   // true
  
- Struct table printing:
  
  type User struct { ID int; Name string }
  _ = utilities.PrintStructTable([]User{{1, "Ada"}, {2, "Linus"}})
  
- Process helpers:
  
  if utilities.DaemonAlreadyRunning("my-app") { /* ... */ }
  pid, err := utilities.FindDaemonProcessPID("my-app")
  
- MIME type by filename:
  
  mime := utilities.MimeTypeFromExtension("photo.JPG") // image/jpeg
  
- Network (ARP):
  
  mac, err := utilities.GetMacAddressFromIp("192.168.1.20")
  
- Gin integration:
  
  func handler(c *gin.Context) {
      claims, err := ginutils.ExtractJwtClaimsFromContext(c, []byte("secret"))
      // ...
  }

## Package Layout

- utilities: general-purpose helpers used across services.
- utilities/ginutils: helpers for extracting JWTs from gin.Context Authorization headers.

## Detailed API (Root package: utilities)

Below is a complete list of exported functions and types with concise descriptions. Examples above show typical usage.

### Pointers and Value Utilities
- func Ptr[T any](v T) *T
  Returns a pointer to v.
- func PtrZeroNil[T any](v T) *T
  Returns &v if v is non-zero; otherwise nil.
- func PtrVal[T any](p *T) T
  Dereferences p or returns zero value if p is nil.
- func PtrCompare[T comparable](p1, p2 *T) bool
  True if both nil or both non-nil and equal.
- func CopyIfNotNil[T any](src, dest *T)
  Copies *src into *dest if both are non-nil.
- func CopyIfNotZero[T any](src T, dest *T)
  Copies src into *dest if src is non-zero and dest != nil.
- func NilIfEmpty[T any](in []T) *[]T
  Returns nil if the slice is empty; otherwise pointer to the slice.
- func NilIfZeroPtr[T comparable](in *T) *T
  Returns nil if in != nil and *in is zero; else returns in.

### Struct Reflection Helpers
- func ZeroStructFieldByName(ptr interface{}, fieldName string) error
  Sets named field to its zero value. ptr must be pointer to struct.
- func SetStructFieldByName(ptr interface{}, fieldName string, value interface{}) error
  Sets named field to value with basic type checking.
- func StructFieldNames(s interface{}) []string
  Returns field names of a struct or pointer to struct.
- func StructToStringMap(s interface{}) map[string]string
  Maps field names to string values; dereferences pointers and marks nil pointers as "<nil>".

### CLI/Table Helpers
- func PrintMapArray(input any) error
  Prints a []map[string]any, []map[string]string, or []utilities.StrMap as a table to stdout.
- func PrintStructMap(obj any) error
  Treats any map's values as structs and prints a table.
- func PrintStructTable(obj any) error
  Prints a struct or slice/array of structs (or pointers) as a table to stdout.

### JWT Helpers
- func GenerateJWT(claims interface{}, duration time.Duration, secretKey []byte) (string, error)
  Creates a JWT (HS256) with iat and exp added.
- func ValidateJWT(tokenString string, secretKey []byte) (*jwt.Token, error)
  Parses and validates an HS256 token.
- func ExtractJwtClaims(tokenString string, secretKey []byte) (map[string]interface{}, error)
  Returns validated claims as a map.
- func ExtractJwtClaimsInto(tokenString string, secretKey []byte, out interface{}) error
  Decodes validated claims into your struct.

### Process/Daemon Helpers
- func DaemonAlreadyRunning(appName string) bool
  True if a matching process already exists.
- func MustFindDaemonProcessPID(appName string) int
  Returns PID if found; returns 0 if not found.
- func FindDaemonProcessPID(appName string) (int, error)
  Cross-platform PID lookup (Linux /proc and macOS variant).
- func FindDaemonProcessPIDWithArg(appName string, argName string) (int, error)
  As above; requires presence of argName in process args.
- func FindProcessPIDMAC(appName string) (int, error)
  macOS-specific lookup using ps.

### Host/Filesystem Helpers
- func AmAdmin() bool
  True if running as root (euid == 0).
- func DirCreateIfNotExists(dir string) error
  mkdir -p behavior with 0755 on missing dirs.

### File Helpers
- func MimeTypeFromExtension(filename string) string
  Returns MIME by extension; defaults to application/octet-stream.

### Network Helpers
- func GetMacAddressFromIp(ipAddress string) (string, error)
  Looks up MAC by IP using local ARP table.

### Database Helpers
- type DbDSNConfig struct { Server string; Port int; Name string; User string; Password string; SSLMode bool; TimeZone string }
  Source values for DSN generation.
- func DbDSN(cfg DbDSNConfig) string
  Assembles a postgres-style DSN from config.
- func ToValuers[T driver.Valuer](in []T) []driver.Valuer
  Useful when building driver.Valuer slices (e.g., for WHERE IN bindings).

### Common Type Aliases
- General maps
  - type StrMap map[string]string
  - type StrAny map[string]any
  - type IntMap map[string]int
  - type AnyMap map[string]any
- JSON
  - type JSON map[string]any
- Slices
  - type Strs []string
  - type Ints []int
  - type Floats []float64
  - type Anys []any
- Functions
  - type Handler func() error
  - type Callback func(any) error
  - type Predicate[T any] func(T) bool
  - type Mapper[T any, R any] func(T) R
- Channels
  - type ErrChan chan error
  - type StrChan chan string
  - type AnyChan chan any

## Subpackage: ginutils

Helpers for extracting JWTs from gin-gonic request contexts.

Import path: github.com/dan-sherwin/go-utilities/ginutils

- func ExtractJwtClaimsFromContext(c *gin.Context, secretKey []byte) (map[string]interface{}, error)
  Reads Authorization: Bearer <token> and returns validated claims.
- func ExtractJwtClaimsFromContextInto(c *gin.Context, secretKey []byte, out interface{}) error
  As above, but decodes into your struct.

Example:

  r := gin.Default()
  r.GET("/me", func(c *gin.Context) {
      var claims struct{ UserID string `json:"uid"` }
      if err := ginutils.ExtractJwtClaimsFromContextInto(c, []byte("secret"), &claims); err != nil {
          c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
          return
      }
      c.JSON(200, gin.H{"user": claims.UserID})
  })

## Notes and Caveats

- Security: JWT helpers use HS256. Ensure secret key management follows your org’s standards. Consider key rotation and short expirations.
- Time: GenerateJWT injects iat and exp based on time.Now().
- Process helpers: Linux uses /proc parsing; macOS uses ps output. Caller’s own PID is ignored. Provide the correct app name (and arg name when applicable).
- Table printing: Helpers write to os.Stdout using tablewriter. In non-TTY contexts, ensure stdout capture is acceptable.
- Network/ARP: GetMacAddressFromIp relies on local ARP cache; may fail if the IP has not been resolved on the local network.
- DSN: DbDSN builds a PostgreSQL-like connection string; adjust as needed for your driver.

## Contributing

- Keep functions small and well-documented (Go doc comments above declarations).
- Prefer generic helpers where it improves ergonomics without obscuring types.
- Add examples to this README when adding new exported functions.

## License

Internal use within Spacelink Corp. Consult repository policy for redistribution.
