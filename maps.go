package utilities

// Merge returns a new map that contains all key-value pairs from a and b.
// If a key exists in both, the value from b takes precedence.
func Merge[K comparable, V any](a, b map[K]V) map[K]V {
	out := make(map[K]V, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

// MergeInto merges all key-value pairs from b into a in place.
// If a key exists in both, the value from b takes precedence.
func MergeInto[K comparable, V any](a, b map[K]V) {
	for k, v := range b {
		a[k] = v
	}
}
