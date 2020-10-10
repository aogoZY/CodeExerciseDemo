package main

import "fmt"

//输入：nums = [1,2,3,4], n = 4, left = 1, right = 5
//输出：13
//解释：所有的子数组和为 1, 3, 6, 10, 2, 5, 9, 3, 7, 4 。
//将它们升序排序后，我们得到新的数组 [1, 2, 3, 3, 4, 5, 6, 7, 9, 10] 。
//下标从 le = 1 到 ri = 5 的和为 1 + 2 + 3 + 3 + 4 = 13 。

func rangeSum(nums []int, n int, left int, right int) int {
	var res []int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		res = append(res, nums[i])
		for j := i + 1; j < len(nums); j++ {
			num += nums[j]
			res = append(res, num)
		}
	}
	fmt.Println(res)
	sortRes := SortList(res)
	fmt.Println(sortRes)
	var countNum int
	for i := left; i <= right; i++ {
		countNum += sortRes[i-1]
	}
	return countNum
}

func SortList(sortList []int) []int {
	for i := len(sortList); i > 0; i-- {
		for j := 0; j < len(sortList)-1; j++ {
			if sortList[j] > sortList[j+1] {
				a := sortList[j]
				sortList[j] = sortList[j+1]
				sortList[j+1] = a
			}
		}
	}
	fmt.Println(sortList)
	return sortList
}

func main() {
	res := rangeSum([]int{1, 2, 3, 4}, 4, 1, 10)
	fmt.Println(res)
}
