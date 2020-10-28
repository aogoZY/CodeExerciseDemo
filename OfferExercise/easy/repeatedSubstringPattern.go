package main

import (
	"fmt"
	"strings"
)

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
//输入: "abab"
//输出: True
//
//输入: "aba"
//输出: False
//
//输入: "abcabcabcabc"
//输出: True


//notice
//1、若是重复的str 则str+str 去掉首尾肯定会包含str
func repeatedSubstringPattern(s string) bool {
	twoStr:=s+s
	//twoStrSplit:=twoStr[1:len(twoStr)-1]
	//fmt.Println(twoStrSplit)
	if strings.Contains(twoStr[1:len(twoStr)-1],s){
		return true
	}
	return false
}

func main() {
	res:=repeatedSubstringPattern("abcabcabcabc")
	fmt.Println(res)
}
