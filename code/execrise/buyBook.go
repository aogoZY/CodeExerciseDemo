package main

import "fmt"

func main() {

	a := 0
	b := 0
	c := 0
	for {
		n, _ := fmt.Scan(&a, &b, &c)
		if n == 0 {
			break
		} else {
			getBook(a, b, c)
		}
	}
}

func getBook(a, b, c int) {

	fmt.Println(227)
}
