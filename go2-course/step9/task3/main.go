package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		res  string
	}{
		{
			name: "test 1",
			in:   "1234",
			res:  "4321",
		},
		{
			name: "test 2",
			in:   "astg",
			res:  "gtsa",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := ReverseString(test.in)
			if ans != test.res {
				t.Fatal("not equal string")
			}
		})
	}
}
