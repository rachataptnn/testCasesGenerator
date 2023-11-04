package main

import (
	"reflect"
	"testing"
)
	
func TestFunc(t *testing.T) {
	type testStruct struct {
		n int
		Output bool
	}

	testcases := []testStruct{
		{
			n: 16,
			Output: true,
		},
		{
			n: 5,
			Output: false,
		},
		{
			n: 1,
			Output: true,
		},
		
	}

	for _, tc := range testcases {
		actual := isPowerOfFour(tc.n, )
		if !reflect.DeepEqual(actual, tc.Output) {
			t.Errorf("Expected %v, but got %v", tc.Output, actual)
		}
	}
}