package main

import "fmt"

//第一站有2个人上车，在第五站有五个人下车，在第三站有三个人上车，在第七站有三个人下车，问在整个过程中是否存在所有乘客树大于容量的情况。
//此时可以考虑用差分方式来做。差分即为后一个数与前一个数的差值，表示其变化。
//[2, 0, 0, 0, -2]		对于第一列他的差分数据为
//[0, 0, 3, 0,  0, 0, -3]	对于第二列差分数据为
//[2, 0, 3, 0,  -2, 0, -3]		即总的差分数据为  所以计算出总的乘客数为[2,2,5,5,3,3,0]。
//所以容量最大为5，输入cap为5时返回true，cap为4时返回false。

func main() {
	input := [][]int{
		{2, 1, 5},
		{3, 3, 7},
	}
	res := carPooling(input, 5)
	fmt.Println(res)
}

func carPooling(trips [][]int, capacity int) bool {
	res := make([]int, 1024)
	var length int
	for j := 0; j < len(trips); j++ {
		pass := trips[j][0]
		start := trips[j][1]
		end := trips[j][2]
		res[start-1] += pass
		res[end-1] -= pass
		if end > length {
			length = end
		}
	}
	countRes := make([]int, length)
	countRes[0] = res[0]
	for i := 1; i < length; i++ {
		countRes[i] = countRes[i-1] + res[i]
		if countRes[i] > capacity {
			return false
		}
	}
	return true
}
