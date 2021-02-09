package main

import "fmt"

func main() {
	res := findContentChildren([]int{1, 2, 3}, []int{1,1})
	fmt.Println(res)
}

func findContentChildren(g []int, s []int) int {
	sortG := sort(g)
	fmt.Println(sortG)
	sortS := sort(s)
	fmt.Println(sortS)
	var res int
	var i, j int
	for i <= len(sortG) -1 && j <= len(sortS)-1 {
		if sortG[i] <= sortS[j] {
			i++
			j++
			res++
		} else {
			j++
		}
	}
	return res
}

func sort(raw []int) (sortedList []int) {
	for i := len(raw) - 1; i > 0; i-- {
		for j := 0; j < len(raw)-1; j++ {
			if raw[j] > raw[j+1] {
				b := raw[j]
				raw[j] = raw[j+1]
				raw[j+1] = b
			}
		}
	}
	return raw
}
