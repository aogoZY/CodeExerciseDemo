package main

import "fmt"

func reverse(x int) int {
	var res int
	for x!=0{
		if tmp :=int32(res);(tmp * 10)  / 10 != tmp{
			return 0
		}
		num := x % 10  //拿到个位数字
		x = x / 10 //拿到除去最后一位的数字
		res = res * 10 + num
	}
	return res
}

func main()  {
	x:=-1230
	res :=reverse(x)
	fmt.Println(res)
}