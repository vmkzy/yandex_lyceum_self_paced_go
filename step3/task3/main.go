package main

func Send(ch1, ch2 chan int) {
	go func() {
		for i := 0; i <= 2; i++ {
			ch1 <- i
		}
	}()
	go func() {
		for i := 0; i <= 2; i++ {
			ch2 <- i
		}
	}()

}
