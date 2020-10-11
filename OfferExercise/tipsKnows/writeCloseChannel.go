package main

import (
	"fmt"
)

//channel关闭之后 写出错 读取决于关闭之前是否有值 有值的话可读取出值 flag为true 否则返回默认值 flag为false

//func main()  {
//	//channel关闭之后，再写不了东西
//	ch:=make(chan string,3)
//	close(ch)
//	ch<-"hello"
//panic: send on closed channel

//channel关闭之后，是否可以读取决于在关闭之前是否还有值，不存在的值为0，flag为false
func main() {
	//int型channel
	cn := make(chan int, 3)
	cn <- 1
	close(cn)
	num1, ok1 := <-cn
	num2, ok2 := <-cn
	num3, ok3 := <-cn
	fmt.Printf("读num channel: %v, %v\n", num1, ok1)
	fmt.Printf("再读num channel: %v, %v\n", num2, ok2)
	fmt.Printf("再再读num channel: %v, %v\n", num3, ok3)

	cs := make(chan string, 3)
	cs <- "a"
	cs <- "b"
	close(cs)
	str1, ok1 := <-cs
	str2, ok2 := <-cs
	str3, ok3 := <-cs
	fmt.Printf("读str channel: %v, %v\n", str1, ok1)
	fmt.Printf("再读str channel: %v, %v\n", str2, ok2)
	fmt.Printf("再再读str channel: %v, %v\n", str3, ok3)

	type Student struct {
		Name string
	}
	cst := make(chan Student, 3)
	cst <- Student{Name: "aogo"}
	close(cst)
	stu1, ok1 := <-cst
	stu2, ok2 := <-cst
	stu3, ok3 := <-cst
	fmt.Printf("读stu channel: %v, %v\n", stu1, ok1)
	fmt.Printf("再读stu channel: %v, %v\n", stu2, ok2)
	fmt.Printf("再再读stu channel: %v, %v\n", stu3, ok3)

//输出：
//	读num channel: 1, true
//	再读num channel: 0, false
//	再再读num channel: 0, false
//	读str channel: a, true
//	再读str channel: b, true
//	再再读str channel: , false
//	读stu channel: {aogo}, true
//	再读stu channel: {}, false
//	再再读stu channel: {}, false


}
