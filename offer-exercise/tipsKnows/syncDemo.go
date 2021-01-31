package main

import (
	"fmt"
	"sync"
)

//func main() {
//	var wg sync.WaitGroup
//	urls := []string{
//		"http://www.golang.org/",
//		"http://www.google.com/",
//		"http://www.somestupidname.com/",
//	}
//	for _, item := range urls {
//		// Increment the WaitGroup counter.
//		wg.Add(1)
//		// Launch a goroutine to fetch the URL.
//		go func(item string) {
//			// Decrement the counter when the goroutine completes.
//			defer wg.Done()
//			// Fetch the URL.
//
//			http.Get(item)
//			fmt.Printf("goroutine working for %s\n",item)
//		}(item)   url 通过 goroutine 的参数进行传递，是为了避免 url 变量通过闭包放入匿名函数后又被修改的问题。
//	}
//	// Wait for all HTTP fetches to complete.
//	wg.Wait()
//	fmt.Println("main is end")
//}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string, 30)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ch chan string, i int) {
			defer wg.Done()
			receiver(ch, i)
		}(ch, i)
	}
	wg.Wait()
	close(ch)
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

func receiver(ch chan string, i int) {
	str:=fmt.Sprintf("this is the %v hello world", i)
	ch <- str
}
