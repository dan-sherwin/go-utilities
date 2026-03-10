package utilities

import (
	"encoding/json"
	"os"
)

// ToJSON converts a value to its JSON string representation.
func ToJSON(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON parses a JSON string into a value.
func FromJSON(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}

// ToJSONIndent converts a value to its indented JSON string representation.
func ToJSONIndent(v any) (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ToJSONFile writes a value to a file in JSON format.
func ToJSONFile(v any, filename string) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// FromJSONFile reads a value from a JSON file.
func FromJSONFile(filename string, v any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// IsJSON checks if a string is a valid JSON.
func IsJSON(data string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(data), &js) == nil
}

// MarshalTo marshals the source value to JSON and then unmarshals it into the destination.
// This is useful for converting between different types that have compatible JSON representations.
func MarshalTo(src any, dst any) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dst)
}
