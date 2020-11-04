package main

import (
	"fmt"
)

//需要将给定数字序列重新排列成字典序中下一个更大的排列。如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//1,2,3 可排列成123、132、213、231、321、312 其中大于且接近123的值为132
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1

//notice:从末尾开始往前找，找到num[i]<num[i+1]的值，并将i+1右边的值最接近num[i]且大于num[i]交换位置，i右边的值反转一波
func main() {
	input := []int{1, 2, 3,4,4}
	//res := reverse(input,1)
	//fmt.Println(res)
	nextPermutation(input)
}

func nextPermutation(input []int) {
	var res []int
	for i := len(input) - 1; i > 0; i-- {
		if input[i-1] < input[i] {
			for j := i; j < len(input)-1; j++ {
				if input[j] >= input[i-1] && input[j+1] <= input[i-1] {
					input[i-1], input[j] = input[j], input[i-1]
					index := i - 1
					res = reverse(input, index)
				}
			}
		}
	}
	fmt.Println(res)
}

func reverse(input []int, index int) []int {
	right := len(input) - 1
	for i := index; i < len(input); i++ {
		if i > right {
			break
		}
		input[i],input[right] =input[right], input[i]
		right--
	}
	return input
}
