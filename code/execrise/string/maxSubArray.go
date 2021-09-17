package main

import (
	"fmt"
	"math"
)

//输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

func main() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}
//暴力解法遍历两边
//高效解法：遍历一遍，每个值累加的过程中，判断累加当前值后sum是否小于0，若是小于0则抛弃前面的，从下一位开始计算子串头。maxValue保存每轮求的的最大值。
func maxSubArray(input []int) int {
	maxValue:=-math.MaxInt32
	var sum int
	for i := 0; i < len(input); i++ {
		sum += input[i]
		if sum < 0 {
			sum = 0
		}

		if sum > maxValue {
			maxValue = sum

		}
	}
	return maxValue
}
