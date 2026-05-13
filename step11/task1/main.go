package main

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Sum[T Number](nums []T) T {
	var res T
	for _, val := range nums {
		res += val
	}
	return res
}
