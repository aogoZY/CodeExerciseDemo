package main

import "fmt"

func findNumberIn2DArray(arr [][]int, target int) bool {
	if len(arr)!=0 && len(arr[0])!=0{
		width := len(arr)
		length := len(arr[0])
		i := 0
		j := length - 1
		for i <= width-1 && j >= 0 {
			if target < arr[i][j] {
				j--
			} else if target > arr[i][j] {
				i++
			} else if target == arr[i][j] {
				return true
			}
		}
	}
	return false
}

func main() {
	Array := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	res := findNumberIn2DArray(Array, 100)
	fmt.Println(res)
}
