package main

import "fmt"

//输入 "1+1"
//输出 2

func main() {
	input := "1+12"
	res := cacluate(input)
	fmt.Println(res)
}

func cacluate(input string) uint8 {
	var res uint8
	var numList []int
	var opera []uint8

	for i := 0; i <= len(input)-1; {
		fmt.Println(input[i])
		var index int
		if input[i] == '+' || input[i] == '-' || input[i] == '*' || input[i] == '/' {
			opera = append(opera, input[i])
		} else {
			var num int
			num =

		}

		i++
	}
	fmt.Println(res)
	return res
}

// + 对应43
// * 对应42
// - 对应45
// / 对应47

// 0 对应 48
