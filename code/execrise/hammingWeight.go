package main

import "fmt"

func hammingWeight(num uint32) int {
	var res int
	for num >0{
		if num & 1 == 1{
			res++
		}
		num = num >>1
	}
	return res
}

func main() {
	res := hammingWeight(00000000000000000000000000001111)
	fmt.Println(res)
}
