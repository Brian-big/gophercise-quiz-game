package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=+=+=+=+=+=+Go Channels=+=+=+=+=+=+")

	channel := make(chan string)

	go count("Lion", channel)
	for msg := range channel {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	fmt.Println("======> Counting", thing, "<======")
	for i := 1; i <= 10; i++ {
		c <- fmt.Sprintf("%d %s",i, thing)
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}
