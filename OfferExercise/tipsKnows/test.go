package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 包含以下基本用法：split、join、for...
func removeDate(s string, char string) string {
	var res []string
	arr := strings.Split(s, "\n")
	for _, val := range arr {
		if strings.Contains(val, char) {
			res = append(res, val)
		}
	}
	// fmt.Println("arr:", arr)
	return strings.Join(res, "\n")
}

func calcSingleDay(input []string, result *map[string]int) {
	for _, val := range input {
		typeOut := strings.Split(val, "-")
		toType := typeOut[0]
		num := typeOut[1]
		if (*result)[toType] != 0 {
			numRes, err := strconv.Atoi(num)
			if err == nil {
				(*result)[toType] += numRes
			}
		} else {
			numRes, err := strconv.Atoi(num)
			if err == nil {
				(*result)[toType] = numRes
			}
		}
	}
	fmt.Println(*result)
}

func calcMoney(input string, result *map[string]int) {
	moneyArr := strings.Split(input, "\n")
	for _, val := range moneyArr {
		costArr := strings.Fields(val)
		calcSingleDay(costArr, result)
	}
}

func main() {
	result := make(map[string]int)
	inputStr := `2020.9.30 周三
 吃饭-126  交通-68  健康-4  日用-100
 
 2020.9.27 周日
 吃饭-358  下馆子-420  交通-5  日用-44  水果-29  学习-149  健康-45`
	moneyStr := removeDate(inputStr, "-")
	fmt.Println(moneyStr)
	calcMoney(moneyStr, &result)
}
