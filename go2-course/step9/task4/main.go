package main

import "testing"

func TestAreAnagrams(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		a    string
		b    string
		res  bool
	}{
		{
			name: "test 1",
			a:    "abc",
			b:    "cba",
			res:  true,
		},
		{
			name: "test 2",
			a:    "a",
			b:    "b",
			res:  false,
		},
		{
			name: "test 3",
			a:    "abab",
			b:    "aba",
			res:  false,
		},
		{
			name: "test 4",
			a:    "AB",
			b:    "ba",
			res:  true,
		},
		{
			name: "test 5",
			a:    "кот",
			b:    "ток",
			res:  true,
		},
		{
			name: "test 6",
			a:    "",
			b:    "",
			res:  true,
		},
		{
			name: "test 7",
			a:    "aab",
			b:    "abb",
			res:  false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ans := AreAnagrams(test.a, test.b)
			if ans != test.res {
				t.Fatalf("error test %s", test.name)
			}
		})
	}
}
