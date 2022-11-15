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

	fmt.Println("+=+=+=+=+= Goroutine select +=+=+=+=+=")

	c1 := make(chan string)
	c2 := make(chan string)

	go func ()  {
		for {
			c1 <- "Channel 1 beeps after 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func ()  {
		for {
			c2 <- "Channel 2 beeps after every 2 seconds"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		fmt.Println(<- c1)
		fmt.Println(<- c2)
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
