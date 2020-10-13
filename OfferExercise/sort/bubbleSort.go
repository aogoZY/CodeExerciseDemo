package main

import "fmt"

func main() {
	input := []int{3, 1, 7, 2, 1, 0, 3}
	res := bubbleSort(input)
	fmt.Println(res)
}

//冒泡排序
//o(n*2)
//两层循环 外层控制处理区间 内层比较相邻的值 大的往后房
//注意 j的取值 避免数组溢出
func bubbleSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
