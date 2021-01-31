package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, _ := reader.ReadLine()

		inputLen := len(input)
		if inputLen <= 0 {
			break
		}
		if inputLen <= 8 {
			fmt.Println("NG")
			continue
		}
		digit, upper, lower, other := 0, 0, 0, 0
		for _, s := range input {
			switch {
			case s >= '0' && s <= '9':
				digit = 1
			case s >= 'a' && s <= 'z':
				lower = 1
			case s >= 'A' && s <= 'Z':
				upper = 1
			default:
				other = 1
			}
		}
		if digit+upper+lower+other < 3 {
			fmt.Println("NG")
			continue
		}
		isRepeat := false
		pw := string(input)
		for i := range pw[:len(pw)-5] {
			if strings.Index(pw[i+3:], pw[i:i+3]) != -1 {
				isRepeat = true
				break
			}
		}
		if isRepeat {
			fmt.Println("NG")
			continue
		}
		fmt.Println("OK")
	}
}
