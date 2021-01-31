package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)


	for {
		oneLine, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		outPut(string(oneLine))
	}
	//outPut("87;A4;W74;W70;A50")
}

func outPut(str string) {
	input := strings.Split(str, ";")

	var a, b int
	for i := 0; i < len(input); i++ {
		a, b = deal(input[i], a, b)
	}

	output := strconv.Itoa(a)+","+strconv.Itoa(b)
	fmt.Println(output)
}

func deal(str string, a, b int) (left, right int) {

	if len(str) > 3 || len(str) < 2 {
		return a, b
	}

	direct := str[0]
	value, err := strconv.Atoi(str[1:])
	if err != nil || (value < 0 || value >= 100) {
		return a, b
	}

	switch direct {
	case 'A':
		return a - value, b
	case 'S':
		return a, b - value
	case 'W':
		return a, b + value
	case 'D':
		return a + value, b
	default:
		return a, b

	}

	return a, b
}