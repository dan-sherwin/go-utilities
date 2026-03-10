package utilities

import (
	"os"
	"testing"
)

func TestToFromJSON(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	input := TestStruct{Name: "John", Age: 30}
	jsonStr, err := ToJSON(input)
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	var output TestStruct
	err = FromJSON(jsonStr, &output)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	if output.Name != input.Name || output.Age != input.Age {
		t.Errorf("FromJSON failed, expected %+v, got %+v", input, output)
	}
}

func TestToJSONIndent(t *testing.T) {
	input := map[string]string{"foo": "bar"}
	jsonStr, err := ToJSONIndent(input)
	if err != nil {
		t.Fatalf("ToJSONIndent failed: %v", err)
	}

	expected := "{\n  \"foo\": \"bar\"\n}"
	if jsonStr != expected {
		t.Errorf("ToJSONIndent failed, expected \n%s\n got \n%s\n", expected, jsonStr)
	}
}

func TestFileJSON(t *testing.T) {
	type TestStruct struct {
		Foo string `json:"foo"`
	}

	tempFile := "test_file_json.json"
	defer func() {
		_ = os.Remove(tempFile)
	}()

	input := TestStruct{Foo: "bar"}
	err := ToJSONFile(input, tempFile)
	if err != nil {
		t.Fatalf("ToJSONFile failed: %v", err)
	}

	var output TestStruct
	err = FromJSONFile(tempFile, &output)
	if err != nil {
		t.Fatalf("FromJSONFile failed: %v", err)
	}

	if output.Foo != input.Foo {
		t.Errorf("File JSON failed, expected %+v, got %+v", input, output)
	}
}

func TestIsJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`{"foo": "bar"}`, true},
		{`[1, 2, 3]`, true},
		{`"hello"`, true},
		{`123`, true},
		{`true`, true},
		{`null`, true},
		{`{invalid}`, false},
		{`hello`, false},
		{``, false},
	}

	for _, tc := range tests {
		if got := IsJSON(tc.input); got != tc.expected {
			t.Errorf("IsJSON(%q) = %v, expected %v", tc.input, got, tc.expected)
		}
	}
}

func TestMarshalTo(t *testing.T) {
	type Src struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}
	type Dst struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	src := Src{Name: "Alice", Age: 25, City: "Paris"}
	var dst Dst

	err := MarshalTo(src, &dst)
	if err != nil {
		t.Fatalf("MarshalTo failed: %v", err)
	}

	if dst.Name != "Alice" || dst.Age != 25 {
		t.Errorf("MarshalTo failed, expected Name=Alice, Age=25, got %+v", dst)
	}
}
