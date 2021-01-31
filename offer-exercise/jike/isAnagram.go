package main

import (
	"fmt"
	"reflect"
)

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//输入: s = "anagram", t = "nagaram"
//输出: true
//
//输入: s = "rat", t = "car"
//输出: false

//1 、将两个字符串排序，判断是否相等
//2、 用map存储该字母出现的次数，比较两个map是否相等  o（n）
func isAnagram(s string, t string) bool {
	return false
}

func main() {
	res := isAnagram2("car", "rat")
	fmt.Println(res)
}

func isAnagram2(str1, str2 string) (res bool) {
	map1 := CountMap(str1)
	map2 := CountMap(str2)
	return reflect.DeepEqual(map1, map2)
}

func CountMap(str string) (res map[int32]int) {
	res = make(map[int32]int)
	for _, item := range str {
		res[item] += 1
	}
	return res
}
