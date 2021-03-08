package main

import "fmt"

//给定两个字符串str1和str2，输出连个字符串的最长公共子序列。如过最长公共子序列为空，则输出-1。
//"1A2C3D4B56","B1D23CA45B6A"
//返回值
//"123456"
//"123456"和“12C4B6”都是最长公共子序列，任意输出一个。

//   12BC23   1A2C3

//小的放后面
func LCS(s1 string, s2 string) string {
	// write code here
	var res string
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}
	var i, j int
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			i++
		} else {
			res += string(s1[i])
			i++
			j++
		}
	}
	return res
}

func main() {
	res := LCS("1A2C3D4B56", "B1D23CA45B6A")
	fmt.Println(res)
}
