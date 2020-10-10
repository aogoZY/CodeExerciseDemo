package main

import (
	"fmt"
)

//实现删除字符串中出现次数最少的字符，若多个字符出现次数一样，则都删除。输出删除这些单词后的字符串，字符串中其它字符保持原来的顺序。
//注意每个输入文件有多组输入，即多个字符串用回车隔开

func main() {
	for {
		var str string
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}
		m := make(map[rune]int)
		for _, j := range str {
			m[j]++
		}
		zx := 20
		for _, val := range m {
			if val < zx {
				zx = val
			}
		}
		var res []rune
		for _, k := range str {
			if m[k] != zx {
				res = append(res, k)
			}
		}
		fmt.Println(string(res))
	}
}
