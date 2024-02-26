package main

import (
	"fmt"
	"time"
)


func StartSender(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- "hello" + name
			time.Sleep(time.Second)
		}
	}()

	return c
}
func main() {
	result := StartSender("An")
	for i := 0; i < 5; i++ {
		fmt.Println(<-result)
	}
}
