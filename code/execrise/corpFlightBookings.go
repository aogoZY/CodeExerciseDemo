package main

import "fmt"

//[[1,2,10],[2,3,20],[2,5,25]], n = 5
//answer = [10,55,45,25,25]

func main() {
	input := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	res := corpFlightBookings2(input, 5)
	fmt.Println(res)
}

func corpFlightBookings(bookings [][]int, n int) []int {
	length := len(bookings)
	countMap := make(map[int]int)
	for j := 0; j < length; j++ {
		start := bookings[j][0]
		end := bookings[j][1]
		for start <= end {
			countMap[start] += bookings[j][2]
			start++
		}
	}
	var res []int
	for i := 1; i <= n; i++ {
		res = append(res, countMap[i])
	}
	return res
}

//差分管理
func corpFlightBookings2(bookings [][]int, n int) []int {
	length := len(bookings)
	res := make([]int, n)

	for j := 0; j < length; j++ {
		start := bookings[j][0]
		end := bookings[j][1]
		for i := start-1; i < end; i++ {
			res[i] += bookings[j][2]
			start++
		}
	}
	return res
}
