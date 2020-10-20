package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	input := []int{1, 7, 4, 3, 1, 8, 5}
	//res := quitSort(input)
	//res := bubbleSortV2(input)
	//res := selectSort(input)
	res := insertSort(input)
	quitSort(input,0,len(input)-1)
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

//冒泡优化
//若已经有序了 冒泡的外循环仍然会遍历所有的元素
//设置一个flag 若已经有序了 则退出循环
func bubbleSortV2(input []int) []int {
	for i := 0; i < len(input); i++ {
		flag := true
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				flag = false
			}

		}
		if flag == true {
			break
		}
	}
	return input
}

//选择排序
//每次找到最小值 放到最小区间的末尾 循环
func selectSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		index := i
		for j := i; j < len(input); j++ {
			if input[j] < input[index] {
				index = j
			}
		}
		input[i], input[index] = input[index], input[i]
	}
	return input
}

//插入排序
//将未排序的元素 和以排序的元素笔记 直到找到以排序元素小于未排序元素 将其插入后面
func insertSort(input []int) []int {
	for i := 1; i < len(input); i++ {
		preIndex := i - 1
		currentValue := input[i]
		for preIndex >= 0 && input[preIndex] > currentValue {
			input[preIndex+1] = input[preIndex]
			preIndex--
		}
		input[preIndex+1] = currentValue
	}
	return input
}

//快排 分成多个子串来做处理 对每个子串 随机选取一个标准值 将小于标准值的数置于左边 将大的标准值置于右边
func quitSort(input []int, left, right int) []int {
	if left < right {
		pos := partion(input, left, right)
		fmt.Println(pos)
		quitSort(input, left, pos-1)
		quitSort(input, pos+1, right)
	}
	return input
}

//返回标志值索引位置
func partion(arr []int, left, right int) int {
	value := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= value {
			arr[i+1], arr[j] = arr[j], arr[i+1]
			i++
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	fmt.Println(arr)
	return i + 1
}

//归并 先对半分 将其子串排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left []int, right [] int) []int {
	leftLength := len(left)
	rightLength := len(right)
	var res []int
	var li, ri int
	for li < leftLength && ri < rightLength {
		if left[li] < right[ri] {
			res = append(res, left[li])
			li++
		} else {
			res = append(res, right[ri])
			ri++
		}
	}

	if li < leftLength {
		res = append(res, left[li:]...)
	} else {
		res = append(res, right[ri:]...)
	}
	return res
}
