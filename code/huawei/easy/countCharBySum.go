package main

import (
	"fmt"
)

func main() {
	available := func(n byte) bool {
		if n == byte(' ') {
			return true
		}
		if n >= byte('0') && n <= byte('9') {
			return true
		}
		if n >= byte('A') && n <= byte('Z') {
			return true
		}
		if n >= byte('a') && n <= byte('z') {
			return true
		}
		return false
	}

	//buff := bufio.NewReader(os.Stdin)
	//for {
		//input, _, err := buff.ReadLine()
		//if err != nil || len(input) == 0 {
		//	return
		//}

		input := "aadddccddc"
		res := make([]int, 256)
		for _, n := range []byte(input) {
			if !available(n) {
				continue
			}
			res[n]++
		}

		for {
			var max int
			for i, c := range res {
				if c == 0 {
					continue
				}
				if res[max] < c {
					max = i
				}
			}
			if res[max] == 0 {
				break
			}
			res[max] = 0
			fmt.Printf("%c", max)
		}
		fmt.Println()
	//}
}
