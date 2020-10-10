package main

import "fmt"

func main() {
	var n int
	var num int
	var rmNum int
	fmt.Scan(&n)
	var typeMap map[int]int
	typeMap = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		typeMap[num]++
	}
	fmt.Scan(&rmNum)
	//fmt.Println(typeMap)
	res := GetRestIphone(n, typeMap, rmNum)
	fmt.Println(res)
}

func GetRestIphone(n int, typeMap map[int]int, rmNum int) (res int) {
	if rmNum == n {
		return 0
	}
	var restAll int
	for item := range typeMap {
		restAll += typeMap[item] - 1
	}
	//fmt.Println("restall",restAll)
	if restAll >= rmNum{
		return len(typeMap)
	}else{
		i :=rmNum - restAll
		res = len(typeMap) - i
		return res
	}
}
