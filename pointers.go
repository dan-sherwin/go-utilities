package utilities

import "reflect"

// Ptr returns a pointer to the given value of a generic type T.
func Ptr[T any](v T) *T { return &v }

// PtrZeroNil returns a pointer to the input value if it is non-zero; otherwise, it returns nil. It uses reflection to determine whether the input value equals its type's zero value.
func PtrZeroNil[T any](v T) *T {
	if reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface()) {
		return nil
	}
	return &v
}

// PtrVal dereferences a pointer and returns its value; if the pointer is nil, it returns the zero value of the type.
func PtrVal[T any](p *T) T {
	var z T
	if p == nil {
		return z
	}
	return *p
}

// PtrCompare compares two pointers of any comparable type and returns true if both are nil or if they point to equal values. Returns false if one pointer is nil or they point to unequal values.
func PtrCompare[T comparable](p1, p2 *T) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}
	return *p1 == *p2
}

// CopyIfNotNil copies the value from the source pointer to the destination pointer only if both pointers are non-nil. It uses PtrVal to dereference the source pointer safely.
func CopyIfNotNil[T any](src, dest *T) {
	if src != nil && dest != nil {
		*dest = PtrVal(src)
	}
}

// CopyIfNotZero copies the value of src to dest if src is not the zero value of its type and dest is not nil. It uses PtrZeroNil to determine if src is non-zero.
func CopyIfNotZero[T any](src T, dest *T) {
	if PtrZeroNil(src) != nil && dest != nil {
		*dest = src
	}
}

// NilIfEmpty takes a slice of any type and returns nil if the slice is empty, or a pointer to the slice otherwise.
func NilIfEmpty[T any](in []T) *[]T {
	if len(in) == 0 {
		return nil
	}
	return &in
}

// NilIfZeroPtr returns nil if the input pointer is non-nil and the value it points to is zero; otherwise, it returns the input pointer unchanged.
func NilIfZeroPtr[T comparable](in *T) *T {
	var zero T
	if in != nil && *in == zero {
		return nil
	}
	return in
}
