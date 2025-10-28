package utilities

type (
	// StrMap is a shorthand for a map of string keys to string values.
	StrMap map[string]string
	// StrAny is a shorthand for a map of string keys to arbitrary values.
	StrAny map[string]any
	// IntMap is a shorthand for a map of string keys to int values.
	IntMap map[string]int
	// AnyMap is a shorthand for a map of string keys to arbitrary values (alias similar to StrAny).
	AnyMap map[string]any

	// JSON is a convenience alias for a generic JSON object represented as a map.
	JSON map[string]any

	// Strs is a shorthand for a slice of strings.
	Strs []string
	// Ints is a shorthand for a slice of ints.
	Ints []int
	// Floats is a shorthand for a slice of float64s.
	Floats []float64
	// Anys is a shorthand for a slice of values of any type.
	Anys []any

	// Handler is a function that returns an error; handy for composing small tasks.
	Handler func() error
	// Callback is a function that receives a value and may return an error.
	Callback func(any) error
	// Predicate is a generic function that returns true/false for a value of type T.
	Predicate[T any] func(T) bool
	// Mapper maps a value of type T to a value of type R.
	Mapper[T any, R any] func(T) R

	// ErrChan is a typed channel for errors.
	ErrChan chan error
	// StrChan is a typed channel for strings.
	StrChan chan string
	// AnyChan is a typed channel for values of any type.
	AnyChan chan any
)
