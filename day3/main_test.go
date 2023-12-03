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
	}

	for _, tt := range tests {
		out := compute(tt.in)
		if out != tt.out {
			t.Errorf("want: %d, got: %d", tt.out, out)
		}
	}
}
