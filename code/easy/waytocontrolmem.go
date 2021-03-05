package main

import (
	"fmt"
	"sync"
	"time"
)

// 控制并发的方式

func main() {
	Count()
	CountByLock()
	CountByWg()
}

func Count() {
	var count1 int
	for i := 0; i < 15000; i++ {
		go func() {
			count1++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count1)
}

//eg2：使用锁+time.wait
func CountByLock() {
	var count2 int
	var mutex sync.Mutex
	for i := 0; i < 15000; i++ {
		go func() {
			defer func() {
				mutex.Unlock()
			}()
			mutex.Lock()
			count2++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count2)
}

//eg3：使用锁+sync.waitgroup
func CountByWg() {
	var count3 int
	var wg sync.WaitGroup
	var mute sync.Mutex
	for i := 0; i < 15000; i++ {
		go func() {
			defer func() {
				wg.Done()
				mute.Unlock()
			}()
			wg.Add(1)
			mute.Lock()
			count3++
		}()
	}
	wg.Wait()
	fmt.Println(count3)
}
