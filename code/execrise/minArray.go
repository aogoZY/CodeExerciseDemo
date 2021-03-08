package main

import "fmt"

func minArray(numbers []int) int {
	smallNum := numbers[0]
	for _, item := range numbers {
		if item < smallNum {
			smallNum = item
		}
	}
	return smallNum
}

func minArray2(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] == numbers[right] {
			right = right - 1

		}
	}
	return numbers[left]
}

func main() {
	res := minArray2([]int{5, 1, 4, 4,4})
	fmt.Println(res)
}
