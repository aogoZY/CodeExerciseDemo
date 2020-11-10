package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for{
			time.Sleep(time.Second)
			func(){
				defer func() {
					if err:=recover();err!=nil{
						fmt.Println("err",err)
					}
				}()
				Pain()
			}()
		}
	}()
}

func HelloFun(ch chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("do sth...")
	ch <- true

}

func Pain() {
	panic("ok")
}
