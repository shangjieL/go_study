package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished...")
}
