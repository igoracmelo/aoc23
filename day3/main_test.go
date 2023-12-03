package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"",
			0,
		},
		{
			".......",
			0,
		},

		// number is at the left
		{
			"123*",
			123,
		},

		// number is at the left
		{
			"33....123*....22..",
			123,
		},

		// number is at the right
		{
			"$403",
			403,
		},

		// number is at the right
		{
			"...432.....$43....123",
			43,
		},

		// numbers on both left and right
		{
			"...42...45$55....123",
			100,
		},
	}

	for i, tt := range tests {
		out := compute(tt.in)
		if out != tt.out {
			t.Errorf("TEST %d FAIL - want: %d, got: %d", i+1, tt.out, out)
		} else {
			t.Logf("TEST %d PASS", i+1)
		}
	}
}
