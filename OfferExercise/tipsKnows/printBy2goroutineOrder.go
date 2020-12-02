package main

import (
	"fmt"
	"time"
)

//起两个协程 分别按顺序打印1-10
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 11; i++ {
			c <- 1
			if i%2 == 0 {
				fmt.Printf("goroutin1 %d\n", i)
			}
		}

	}()

	go func() {
		for i := 0; i < 11; i++ {
			<-c
			if i%2 == 1 {
				fmt.Printf("goroutin2 %d\n", i)

			}
		}
	}()

	time.Sleep(time.Second * 5)
}
