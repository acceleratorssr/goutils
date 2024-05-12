package _map

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

type TestStructPtr struct {
	Name     *string `json:"name"`
	Age      *int    `json:"age"`
	Location *string `json:"location"`
}

func assertEqual(t *testing.T, expected, actual map[interface{}]interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("测试失败：期望值 %v，实际值 %v", expected, actual)
	}
}

func TestStruct2Map(t *testing.T) {
	name := "acc"
	age := 30
	location := "GuangZhou"

	testStruct := TestStruct{
		Name:     name,
		Age:      age,
		Location: location,
	}

	testStructPtr := TestStructPtr{
		Name:     &name,
		Age:      &age,
		Location: &location,
	}

	expectedCompleteMap := map[interface{}]interface{}{
		"Name":     name,
		"Age":      age,
		"Location": location,
	}

	tests := []struct {
		name     string
		input    interface{}
		fields   []string
		expected map[interface{}]interface{}
	}{
		{
			name:     "TestBasicConversion",
			input:    testStruct,
			fields:   []string{"Name", "Age", "Location"},
			expected: expectedCompleteMap,
		},
		{
			name:     "TestEmptyStruct",
			input:    TestStruct{},
			fields:   []string{"Name", "Age", "Location"},
			expected: map[interface{}]interface{}{},
		},
		{
			name:     "TestMissingFields",
			input:    testStruct,
			fields:   []string{"Name", "MissingField"},
			expected: map[interface{}]interface{}{"Name": name},
		},
		{
			name:     "TestFieldsPtr",
			input:    testStructPtr,
			fields:   []string{"Name", "Age", "Location"},
			expected: expectedCompleteMap,
		},
		{
			name:     "TestStructPtr",
			input:    &testStruct,
			fields:   []string{"Name", "Age", "Location"},
			expected: expectedCompleteMap,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Struct2Map(tt.input, tt.fields)
			assertEqual(t, tt.expected, result)
		})
	}
}
