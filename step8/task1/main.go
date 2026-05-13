package main

import (
	"math"
	"time"
)

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	time.AfterFunc(100*time.Millisecond, func() {
		select {
		case <-stop:
			return
		default:
			close(stop)
		}
	})
	defer close(prime_nums)
	for i := 2; i <= N; i++ {
		select {
		case <-stop:
			return
		default:
		}
		sq := int(math.Sqrt(float64(i)))
		f := true
		for j := 2; j <= sq; j++ {
			if i%j == 0 {
				f = false
				break
			}
		}
		if f {
			select {
			case <-stop:
				return
			case prime_nums <- i:
			}
		}
	}
}
