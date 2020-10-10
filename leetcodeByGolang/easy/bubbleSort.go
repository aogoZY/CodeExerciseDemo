package main

import "fmt"

func main() {
	input := []int{3, 2, 6, 7, 1, 9,8,0}
	bubble(input)
}

func bubble(input []int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	fmt.Println(input)
}
