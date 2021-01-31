package main

import "fmt"

func fib(n int) int {
	var res map[int]int
	res = make(map[int]int)
	res[0] = 0
	res[1] = 1
	if n >= 2 {
		for i := 2; i <=n; i++ {
			res[i] = (res[i-1] + res[i-2]) %  1000000007
		}
		return res[n]
	}
	return res[n]
}

func main() {
	res := fib(95)
	fmt.Println(res)
}
