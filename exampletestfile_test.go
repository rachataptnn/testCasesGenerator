package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	type testStruct struct {
		n      int
		pick   int
		Output int
	}

	testcases := []testStruct{
		{
			n:      10,
			pick:   6,
			Output: 6,
		},
		{
			n:      1,
			pick:   1,
			Output: 1,
		},
		{
			n:      1,
			pick:   1,
			Output: 1,
		},
	}

	for _, tc := range testcases {
		actual := leetcodefuncName(tc.n, tc.pick)
		if !reflect.DeepEqual(actual, tc.Output) {
			t.Errorf("Expected %v, but got %v", tc.Output, actual)
		}
	}
}
