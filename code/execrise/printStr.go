package main

import (
	"fmt"
	sort2 "sort"
)

func main() {

	a := 0
	var input []string
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {

			for a > 0 {
				var str string
				fmt.Scan(&str)
				input = append(input, str)
				a--
				//fmt.Println(input)
			}
			fmt.Println(input)

			GetOrder(input)
		}
	}
}

//func main() {
//	a := "abc"
//	b := "123456789"
//	GetOrder(a, b)
//}

func GetOrder(a []string) {
	var allRes []string
	for _, item := range a {
		allRes = append(allRes, GetA(item)...)
	}
	sort2.Strings(allRes)
	for _, item := range allRes {
		fmt.Print(item)
		fmt.Print(" ")
	}
}

func GetA(a string) []string {
	newA := makeLength(a)
	resA := SplitStr(newA)
	return resA
}

func makeLength(str string) string {
	length := len(str)
	beishu := length/8 + 1
	for i := 0; i < beishu*8-length; i++ {
		str = str + "0"
	}
	return str
}

func SplitStr(str string) []string {
	var res []string
	for i := 0; i < len(str); i++ {
		if i%8 == 0 {
			res = append(res, str[i:i+8])
		}
	}
	return res
}
