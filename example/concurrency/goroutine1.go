package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(1)
	go sender(ch)
	go consumer(ch, &wg)
	wg.Wait()
}

func sender(ch chan<- string) {
	ch <- "flower"
	ch <- "bird"
	ch <- "duck"
	ch <- "groff"
	ch <- "people"
}

func consumer(ch <-chan string, wg *sync.WaitGroup) {
	num := 0
	for num < 5 {
		data := <-ch
		if data != "" {
			num++
		}
		fmt.Printf("Received data: %s\n", data)
	}
	wg.Done()
}
