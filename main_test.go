Example 1:
Input: n = 10, pick = 6
Output: 6
Example 2:
Input: n = 1, pick = 1
Output: 1
Example 3:
Input: n = 2, pick = 1
Output: 1


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

