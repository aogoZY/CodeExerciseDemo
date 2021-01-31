package main

import (
	"fmt"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

//思路：
//1、暴力 a、b、c 3层循环  o(n*n*n)
//2、 c= -(a+b)  o(n*n) a、b 2层循环
//3、 先排序，a循环，在子树组利用两边向中间夹的方式，判断a+b+>0;最右边往左移，<0 最左边往右移
func threeSum(nums []int) [][]int {
	var resList [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				var resItem []int
				if nums[i]+nums[j]+nums[k] == 0 {
					resItem = append(resItem, nums[i], nums[j], nums[k])
					resList = append(resList, resItem)
				}
			}
		}
	}
	return resList
}

func threeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	//先将数组排序，若最小的值已经大于0，直接跳出循环，因为后面的数相加一定>0
	sort.Ints(nums)

	var resList [][]int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			break
		}
		//避免有重复值，若外层循环的值相等，则跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				resList = append(resList, []int{nums[i], nums[left], nums[right]})
				//若当前值已经相等，则需要left和right都往里走一个位，直到没有重复的l/r
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}

			} else if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			}
		}

	}
	return resList
}

func main() {
	res := threeSum2([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
