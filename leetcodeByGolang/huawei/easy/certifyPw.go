package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//密码要求:
//1.长度超过8位
//2.包括大小写字母.数字.其它符号,以上四种至少三种
//3.不能有相同长度大于2的子串重复

//021Abc9000
//021Abc9Abc1
//021ABC9000
//021$bc9000

//OK
//NG
//NG
//OK

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
	str, _, err := inputReader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	if len(str) < 8 {
		fmt.Println("NG")
	}
	//str := "021Abc9Abc1"
	var smallChar, bigChar, num, other int
	for _, item := range str {
		switch {
		case item <= 'z' && item >= 'a':
			smallChar = 1
		case item <= 'Z' && item >= 'A':
			bigChar = 1
		case item <= '9' && item >= '0':
			num = 1
		default:
			other = 1
		}
	}
	if smallChar+bigChar+num+other < 3 {
		fmt.Println("NG")
		continue
	}
	var isRepeat bool
		//for i := range str[:len(str)-5] {

	for i := 0; i < len(str)-5; i++ {
		targrt := string(str[i:i+3])
		subStr := string(str[i+3:])
		res := strings.Index(subStr, targrt)
		if res == -1 {
			isRepeat = true
			break
		}
	}
	if !isRepeat {
		fmt.Println("OK")

	} else {
		fmt.Println("NG")
	}
	}

}
