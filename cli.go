package utilities

import (
	"fmt"
	"os"
	"reflect"
	"sort"

	"github.com/olekukonko/tablewriter"
)

// PrintMapArray takes an input of type any, attempts to interpret it as a slice of maps with string keys and values of any type, and prints it as a formatted table. Returns an error if the input is of unsupported type or is an empty slice.
func PrintMapArray(input any) error {
	var ma []map[string]any

	switch v := input.(type) {
	case []map[string]any:
		ma = v
	case []map[string]string:
		for _, m := range v {
			converted := make(map[string]any, len(m))
			for k, val := range m {
				converted[k] = val
			}
			ma = append(ma, converted)
		}
	case []StrMap:
		for _, m := range v {
			converted := make(map[string]any, len(m))
			for k, val := range m {
				converted[k] = val
			}
			ma = append(ma, converted)
		}
	default:
		return fmt.Errorf("unsupported type for PrintMapArray")
	}

	if len(ma) == 0 {
		return fmt.Errorf("input slice is empty")
	}

	table := tablewriter.NewWriter(os.Stdout)
	header := []string{}
	for k := range ma[0] {
		header = append(header, k)
	}
	table.Header(header)
	for _, m := range ma {
		values := []string{}
		for _, k := range header {
			values = append(values, fmt.Sprintf("%v", m[k]))
		}
		if err := table.Append(values); err != nil {
			return err
		}
	}
	if err := table.Render(); err != nil {
		return err
	}
	return nil
}

// PrintStructMap takes any map as input and prints its values in a tabular format by treating them as structs.
// Returns an error if the input is not a map or if any issues occur during processing.
func PrintStructMap(obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Map {
		return fmt.Errorf("input must be a map")
	}

	values := make([]interface{}, 0, v.Len())
	for iter := v.MapRange(); iter.Next(); {
		values = append(values, iter.Value().Interface())
	}
	return PrintStructTable(values)
}

// PrintSortedStructMap takes any map as input and prints its values in a tabular format by treating them as structs.
// It behaves like PrintStructMap but sorts the output by the map key before printing.
func PrintSortedStructMap(obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Map {
		return fmt.Errorf("input must be a map")
	}

	if v.Len() == 0 {
		return nil
	}

	keys := v.MapKeys()
	keyKind := v.Type().Key().Kind()

	switch keyKind {
	case reflect.String:
		sort.Slice(keys, func(i, j int) bool { return keys[i].String() < keys[j].String() })
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		sort.Slice(keys, func(i, j int) bool { return keys[i].Int() < keys[j].Int() })
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		sort.Slice(keys, func(i, j int) bool { return keys[i].Uint() < keys[j].Uint() })
	case reflect.Float32, reflect.Float64:
		sort.Slice(keys, func(i, j int) bool { return keys[i].Float() < keys[j].Float() })
	default:
		// Fallback: compare string representations
		// Note: This ensures deterministic order even for unsupported key kinds
		sort.Slice(keys, func(i, j int) bool {
			return fmt.Sprint(keys[i].Interface()) < fmt.Sprint(keys[j].Interface())
		})
	}

	values := make([]interface{}, 0, v.Len())
	for _, k := range keys {
		values = append(values, v.MapIndex(k).Interface())
	}
	return PrintStructTable(values)
}

// PrintStructTable prints a tabular representation of a struct or a slice/array of structs to the standard output.
// It requires the input to be a struct, or a slice/array containing structs or pointers to structs.
// Returns an error if input is invalid or processing fails.
func PrintStructTable(obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		objSlice := reflect.MakeSlice(reflect.SliceOf(v.Type()), 0, 1)
		objSlice = reflect.Append(objSlice, v)
		v = objSlice
	case reflect.Slice, reflect.Array:
		vLen := v.Len()
		if vLen == 0 {
			return nil
		}
		if v.Len() < 1 || (v.Index(0).Kind() != reflect.Struct && v.Index(0).Kind() != reflect.Ptr && v.Index(0).Kind() != reflect.Interface) {
			return fmt.Errorf("input slice/array must contain structs or pointers to structs")
		}
	default:
		return fmt.Errorf("input must be a struct or slice/array of structs")
	}

	if v.Len() == 0 {
		return fmt.Errorf("input slice/array is empty")
	}

	value := v.Index(0)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	fieldNames := StructFieldNames(value.Interface())
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(fieldNames)
	for i := 0; i < v.Len(); i++ {
		value = v.Index(i)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		stringMap := StructToStringMap(value.Interface())
		tableValues := []string{}
		for _, fieldName := range fieldNames {
			tableValues = append(tableValues, stringMap[fieldName])
		}
		if err := table.Append(tableValues); err != nil {
			return err
		}
	}
	if err := table.Render(); err != nil {
		return err
	}
	return nil
}

// Additional table writer helpers for simple slices and maps

// PrintStringSlice prints a one-column table of a []string or utilities.Strs with row numbers.
// Returns an error if the input slice is empty or of unsupported type.
func PrintStringSlice(input any) error {
	var arr []string
	switch v := input.(type) {
	case []string:
		arr = v
	case Strs:
		arr = []string(v)
	default:
		return fmt.Errorf("unsupported type for PrintStringSlice")
	}
	if len(arr) == 0 {
		return fmt.Errorf("input slice is empty")
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"#", "Value"})
	for i, s := range arr {
		if err := table.Append([]string{fmt.Sprintf("%d", i), s}); err != nil {
			return err
		}
	}
	if err := table.Render(); err != nil {
		return err
	}
	return nil
}

// PrintAnySlice prints a one-column table of a []any or utilities.Anys with row numbers.
// Values are stringified using fmt.Sprintf("%v", v).
func PrintAnySlice(input any) error {
	var arr []any
	switch v := input.(type) {
	case []any:
		arr = v
	case Anys:
		arr = []any(v)
	default:
		return fmt.Errorf("unsupported type for PrintAnySlice")
	}
	if len(arr) == 0 {
		return fmt.Errorf("input slice is empty")
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"#", "Value"})
	for i, val := range arr {
		if err := table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%v", val)}); err != nil {
			return err
		}
	}
	if err := table.Render(); err != nil {
		return err
	}
	return nil
}

// PrintMap prints a two-column table for any map type.
// It accepts maps with any key and value types (including pointers to maps).
// Keys are ordered by their string representation for stable output.
// Optional headers can be provided: headers[0] for the key column, headers[1] for the value column.
func PrintMap(input any, headers ...string) error {
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Map {
		return fmt.Errorf("input must be a map")
	}
	if !v.IsValid() || v.Len() == 0 {
		return fmt.Errorf("input map is empty")
	}

	// Collect keys and sort them by their string representation for deterministic output
	keys := v.MapKeys()
	type kv struct {
		k    reflect.Value
		kStr string
	}
	pairs := make([]kv, 0, len(keys))
	for _, k := range keys {
		pairs = append(pairs, kv{k: k, kStr: fmt.Sprintf("%v", k.Interface())})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].kStr < pairs[j].kStr })

	keyHeader := "Key"
	valueHeader := "Value"
	if len(headers) >= 1 && headers[0] != "" {
		keyHeader = headers[0]
	}
	if len(headers) >= 2 && headers[1] != "" {
		valueHeader = headers[1]
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{keyHeader, valueHeader})
	for _, p := range pairs {
		val := v.MapIndex(p.k)
		if err := table.Append([]string{p.kStr, fmt.Sprintf("%v", val.Interface())}); err != nil {
			return err
		}
	}
	if err := table.Render(); err != nil {
		return err
	}
	return nil
}

// PrintStringsTable prints a table for a [][]string with optional headers.
// If headers is nil or empty, rows are printed without a header.
// Returns an error if rows is empty or contains inconsistent column counts when headers are provided.
func PrintStringsTable(headers []string, rows [][]string) error {
	if len(rows) == 0 {
		return fmt.Errorf("input rows are empty")
	}
	if len(headers) > 0 {
		for i := range rows {
			if len(rows[i]) != len(headers) {
				return fmt.Errorf("row %d has %d columns, expected %d", i, len(rows[i]), len(headers))
			}
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	if len(headers) > 0 {
		table.Header(headers)
	}
	for _, r := range rows {
		if err := table.Append(r); err != nil {
			return err
		}
	}
	if err := table.Render(); err != nil {
		return err
	}
	return nil
}

// PrintSlice prints any slice or array (of basic types or structs) as a two-column table of index and value.
// For struct elements, it will use fmt.Sprintf("%v", elem). For slices of structs requiring field breakdown, use PrintStructTable instead.
func PrintSlice(input any) error {
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return fmt.Errorf("input must be a slice or array")
	}
	if v.Len() == 0 {
		return fmt.Errorf("input slice/array is empty")
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"#", "Value"})
	for i := 0; i < v.Len(); i++ {
		if err := table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%v", v.Index(i).Interface())}); err != nil {
			return err
		}
	}
	if err := table.Render(); err != nil {
		return err
	}
	return nil
}
