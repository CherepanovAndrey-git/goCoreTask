package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name             string
		a                []int
		b                []int
		expectedExist    bool
		expectedElements []int
	}{
		{
			name:             "basic example",
			a:                []int{65, 3, 58, 678, 64},
			b:                []int{64, 2, 3, 43},
			expectedExist:    true,
			expectedElements: []int{3, 64},
		},
		{
			name:             "no intersection",
			a:                []int{1, 2, 3},
			b:                []int{4, 5, 6},
			expectedExist:    false,
			expectedElements: []int{},
		},
		{
			name:             "duplicates in first slice",
			a:                []int{3, 3, 64},
			b:                []int{3, 64},
			expectedExist:    true,
			expectedElements: []int{3, 64},
		},
		{
			name:             "empty first slice",
			a:                []int{},
			b:                []int{1, 2},
			expectedExist:    false,
			expectedElements: []int{},
		},
		{
			name:             "empty second slice",
			a:                []int{1, 2},
			b:                []int{},
			expectedExist:    false,
			expectedElements: []int{},
		},
		{
			name:             "full intersection",
			a:                []int{5, 5, 5},
			b:                []int{5},
			expectedExist:    true,
			expectedElements: []int{5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exist, result := FindIntersection(tt.a, tt.b)

			if exist != tt.expectedExist {
				t.Errorf("Exist: got %v, want %v", exist, tt.expectedExist)
			}

			sort.Ints(result)
			sort.Ints(tt.expectedElements)
			
			if !reflect.DeepEqual(result, tt.expectedElements) {
				t.Errorf("Elements: got %v, want %v", result, tt.expectedElements)
			}
		})
	}
}
