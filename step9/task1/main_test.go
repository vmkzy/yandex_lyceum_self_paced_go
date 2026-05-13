package main

import (
	"reflect"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		res  []int
	}{
		{
			name: "ok",
			in:   []int{1, 2, 3, 4},
			res:  []int{1, 2, 3, 4},
		},
		{
			name: "ne ok",
			in:   []int{4, 2, 6},
			res:  []int{2, 4, 6},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SortIntegers(test.in)
			if !reflect.DeepEqual(test.in, test.res) {
				t.Fatal("no equal")
			}
		})
	}
}
