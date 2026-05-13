package main

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		in     []int
		target int
		res    bool
	}{
		{
			name:   "1",
			in:     []int{1, 2, 3, 4},
			target: 1,
			res:    true,
		},
		{
			name:   "2",
			in:     []int{1, 2, 3, 4},
			target: 5,
			res:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := Contains(test.in, test.target)
			if ans != test.res {
				t.Fatal("error")
			}
		})
	}
}
