package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//输出字符串中最长的数字字符串和它的长度，中间用逗号间隔。
// 如果有相同长度的串，则要一块儿输出（中间不间隔），但是长度还是一串的长度，与数字字符串间用逗号间隔。

//abcd12345ed125ss123058789
//123058789,9

//思路1：遍历字符串，对于数字类型用temp变量记录长度，取当前str【i-temp+1：i】为数字字符子串。
//若下一个子数字子串出现，比较其与当前maxLength长度，长则置换。相等取其res+=str【i-temp+1：i】

//思路2：将不是数字的值全部换成'A'，再用'A'split字符串。

func main() {
	//var str string
	//str = "ad12345ed125ss123058789"
	//fmt.Scan(&str)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		res, length := GetLongNumAndLength2(input.Text())
		lengthStr := strconv.Itoa(length)
		outPut := res + "," + lengthStr
		fmt.Print(outPut)
	}

}

func GetLongNumAndLength(str string) (res string, length int) {
	if len(str) < 1 {
		return
	}
	var temp int
	var max int
	for i := 0; i < len(str); i++ {
		//fmt.Println(str[i])
		if str[i] < '0' || str[i] > '9' {
			temp = 0
		} else {
			temp++
			if temp > max {
				max = temp
				res = str[i-temp+1 : i+1]
			} else if temp == max {
				res += str[i-temp+1 : i+1]
			}
		}
	}
	return res, max
}

func GetLongNumAndLength2(str string) (res string, length int) {
	var fmtStr string
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			fmtStr += "A"
		} else {
			fmtStr += string(str[i])
		}
	}
	//fmt.Println(fmtStr)
	splitRes := strings.Split(fmtStr, "A")
	max := 0
	for i := 0; i < len(splitRes); i++ {
		if max < len(splitRes[i]) {
			max = len(splitRes[i])
		}
	}
	for i := 0; i < len(splitRes); i++ {
		if len(splitRes[i]) == max {
			res += splitRes[i]
		}
	}

	return res, max
}
