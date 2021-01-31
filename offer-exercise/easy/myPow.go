package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n > 0 {
		return Multipy(x, n)
	} else {
		res :=Multipy(x, -n)
		return 1 /res
	}
}

func Multipy(x float64, n int) (res float64) {
	res = 1.0
	for n > 0 {
		res = res * x
		n--
	}
	return res
}

func main() {
	res := myPow(2.0, -2)
	fmt.Println(res)
}
