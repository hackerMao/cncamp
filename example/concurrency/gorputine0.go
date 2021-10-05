package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan<- string) {
	ch <- "flower"
	ch <- "bird"
	ch <- "duck"
	ch <- "grass"
	ch <- "people"
}

func getData(ch <-chan string) {
	for {
		data := <-ch
		fmt.Printf("Received data: %s\n", data)
	}
}
