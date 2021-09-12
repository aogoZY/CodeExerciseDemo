package main

import "fmt"

func main(){
	res := findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
	fmt.Println(res)
}

func findRepeatNumber(nums []int) int {
	for i,item :=range nums{
		for i:=i+1;i<len(nums);i++{
			if nums[i]==item{
				return item
			}
		}
	}
	return 0
}