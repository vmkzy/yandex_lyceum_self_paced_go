package main

import (
	"errors"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	if n < 0 {
		return 0, errors.New("n must be non-negative")
	}
	res := make(chan int, 1)

	go func() {
		if n == 0 {
			res <- 0
		}
		if n == 1 {
			res <- 1
		}
		prev, cur := 0, 1
		for i := 2; i <= n; i++ {
			prev, cur = cur, prev+cur
		}
		res <- cur
	}()

	select {
	case ans := <-res:
		return ans, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}

}
