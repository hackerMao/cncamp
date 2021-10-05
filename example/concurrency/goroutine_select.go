package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump1(ch2)
	go suck(ch1, ch2)

	time.Sleep(time.Second)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("Received data from channel 1, data: %d\n", v1)
		case v2 := <-ch2:
			fmt.Printf("Received data from channel 2, data: %d\n", v2)

		}
	}
}
