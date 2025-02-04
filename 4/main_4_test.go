package main

import (
	"reflect"
	"testing"
)

func TestSliceDiff(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "basic example",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "no common elements",
			slice1:   []string{"a", "b"},
			slice2:   []string{"c", "d"},
			expected: []string{"a", "b"},
		},
		{
			name:     "all elements from slice1 in slice2",
			slice1:   []string{"a", "b"},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "empty slice1",
			slice1:   []string{},
			slice2:   []string{"a"},
			expected: []string{},
		},
		{
			name:     "empty slice2",
			slice1:   []string{"a"},
			slice2:   []string{},
			expected: []string{"a"},
		},
		{
			name:     "duplicates in slice1",
			slice1:   []string{"a", "a", "b"},
			slice2:   []string{"b"},
			expected: []string{"a", "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceDiff(tt.slice1, tt.slice2)

			if result == nil {
				result = []string{}
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("For test '%s' got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}
