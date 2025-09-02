package utilities

type (
	// General map shortcuts
	StrMap map[string]string
	StrAny map[string]any
	IntMap map[string]int
	AnyMap map[string]any // Same as StrAny, just alternate naming

	// JSON helpers
	JSON map[string]any // explicitly for JSON payloads

	// Slice shortcuts
	Strs   []string
	Ints   []int
	Floats []float64
	Anys   []any

	// Function aliases
	Handler              func() error
	Callback             func(any) error
	Predicate[T any]     func(T) bool
	Mapper[T any, R any] func(T) R

	// Channel aliases
	ErrChan chan error
	StrChan chan string
	AnyChan chan any
)
