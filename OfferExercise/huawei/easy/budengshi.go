package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	res := strings.Split(str, ";")
	length := len(res)
	//a1 := res[0]
	//a2 := res[1]
	//a3 := res[2]
	x := res[length-3]
	b := res[length-2]
	char := res[length-1]

	//a1List := strings.Split(a1, ",")
	//a2List := strings.Split(a2, ",")
	//a3List := strings.Split(a3, ",")
	xList := strings.Split(x, ",")
	bList := strings.Split(b, ",")
	charList := strings.Split(char, ",")
	flagEnd := true
	var maxResult float64
	for i := 0; i < 3; i++ {
		bValue, _ := strconv.ParseFloat(bList[i], 64)
		flag, count := IsTrue(strings.Split(res[i], ","), xList, bValue, charList[i])
		if count > maxResult {
			maxResult = count
		}
		if flagEnd {
			if !flag {
				flagEnd = false
			}
		}
	}
	maxResultStr := strconv.FormatFloat(maxResult, 'g', -1, 64)
	max := strings.Split(maxResultStr, ".")
	maxInt,_ := strconv.Atoi(max[0])
	fmt.Println(flagEnd, maxInt)

}

func IsTrue(aList []string, xList []string, b float64, char string) (bool, float64) {
	var resA1 float64
	for i := 0; i < len(aList); i++ {
		a1Float, _ := strconv.ParseFloat(aList[i], 64)
		xFloat, _ := strconv.ParseFloat(xList[i], 64)
		resA1 += a1Float * xFloat
	}
	res := resA1 - b
	if b <= resA1 && char == "<=" {
		return true, res
	}
	if b >= resA1 && char == ">=" {
		return true, res
	}
	return false, res
}
