package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//A表示向左移动，D表示向右移动，W表示向上移动，S表示向下移动。

func PointMove(str []string) (x int, y int) {
	for _, item := range str {
		num, _ := strconv.Atoi(item[1:])
		if strings.HasPrefix(item, "A") {
			x -= num
		} else if strings.HasPrefix(item, "D") {
			x += num
		} else if strings.HasPrefix(item, "W") {
			y += num
		} else if strings.HasPrefix(item, "S") {
			y -= num
		}
	}
	return x, y
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		str = str[:len(str)-1]

		getVerfityStr(string(str))
	}
}

func getVerfityStr(str string) {
	splitRes := strings.Split(str, ";")
	var strList []string
	for _, item := range splitRes {
		matchFlag := match(item)
		if matchFlag {
			strList = append(strList, item)
		}
	}
	x, y := PointMove(strList)
	output := strconv.Itoa(x) + "," + strconv.Itoa(y)
	fmt.Println(output)

}

func match(str string) bool {
	if str == "" || len(str) == 0 || len(str) > 4 || len(str) < 2 {
		return false
	}

	if !unicode.IsLetter(rune(str[0])) {
		return false
	}
	strInt, err := strconv.Atoi(str[1:])
	if err != nil {
		return false
	}
	if strInt > 0 && strInt < 100 {
		return true
	}
	return false
}
