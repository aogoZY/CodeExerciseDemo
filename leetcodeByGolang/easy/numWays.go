package main

import "fmt"

func numWays(n int) int {
	if n== 1 || n==2{
		return n
	}
	if n==0{
		return 1
	}
	var res map[int]int
	res =make(map[int]int)
	for i:=3;i<=n;i++{
		res[1]=1
		res[2]=2
		res[i] = (res[i-2]+res[i-1]) % 1000000007
	}
	return res[n]
}

func main()  {
	res := numWays(7)
	fmt.Println(res)
}