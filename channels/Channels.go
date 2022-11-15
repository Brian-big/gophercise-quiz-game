package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=+=+=+=+=+=+Go Channels=+=+=+=+=+=+")

	go func() {
		count("Lion")
	}()
	go func ()  {
		count("Mike")
	}()
	count("Sheep")

}

func count(thing string) {
	fmt.Println("======> Counting", thing, "<======")
	for i := 1; i <= 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(500 * time.Millisecond)
	}
}
