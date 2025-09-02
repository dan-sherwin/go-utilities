package utilities

import (
	"fmt"
	"reflect"
)

// ZeroStructFieldByName sets the specified field of a struct to its zero value.
// The `ptr` parameter must be a pointer to a struct, and `fieldName` is the name of the field to reset.
// Returns an error if the field does not exist, cannot be set, or `ptr` is not a pointer to a struct.
func ZeroStructFieldByName(ptr interface{}, fieldName string) error {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}
	v = v.Elem()
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s in struct", fieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %s in struct", fieldName)
	}
	field.Set(reflect.Zero(field.Type()))
	return nil
}

// SetStructFieldByName sets the value of a field in a struct identified by its name.
// The function expects a pointer to a struct, the field name as a string, and the value to set.
// It returns an error if the pointer is not to a struct, the field does not exist or cannot be set,
// or if there is a type mismatch between the field and the value.
func SetStructFieldByName(ptr interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(ptr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}

	v = v.Elem()
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s in struct", fieldName)
	}
	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %s in struct", fieldName)
	}

	val := reflect.ValueOf(value)
	if field.Kind() != val.Kind() {
		return fmt.Errorf("provided value type does not match field type")
	}

	field.Set(val)
	return nil
}

// StructFieldNames returns a slice of field names for a given struct or a pointer to a struct.
// If the input is not a struct or a pointer to a struct, it returns an empty slice.
func StructFieldNames(s interface{}) []string {
	fields := []string{}
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			fields = append(fields, t.Field(i).Name)
		}
	}
	return fields
}

// StructToStringMap converts a struct into a map with field names as keys and field values as strings.
// It supports nested pointers by dereferencing them and handles nil pointers, assigning "<nil>" as their value.
// The input must be a struct or a pointer to a struct; otherwise, behavior is undefined.
func StructToStringMap(s interface{}) map[string]string {
	m := make(map[string]string)
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		rval := reflect.ValueOf(v.Field(i).Interface())
		var valueStr string
		if rval.Kind() == reflect.Ptr {
			if !rval.IsNil() {
				valueStr = fmt.Sprintf("%v", rval.Elem().Interface())
			} else {
				valueStr = "<nil>"
			}
		} else {
			valueStr = fmt.Sprintf("%v", v.Field(i).Interface())
		}
		m[t.Field(i).Name] = valueStr
	}
	return m
}
