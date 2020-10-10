package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	return addDigits(GetSum(num))
}

func GetSum(num int) int {
	var rest int = num
	var count int
	for rest > 0 {
		weishu := rest % 10
		rest = rest / 10
		count += weishu
	}
	return count

}

func main() {
	res := addDigits(199)
	fmt.Println(res)
}
