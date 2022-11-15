package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=+=+=+=+=+=+Go Channels=+=+=+=+=+=+")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("Lion")
		wg.Done()
	}()

	wg.Wait()

}

func count(thing string) {
	fmt.Println("======> Counting", thing, "<======")
	for i := 1; i <= 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(500 * time.Millisecond)
	}
}
