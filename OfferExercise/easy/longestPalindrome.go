package main

import "fmt"

//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。请注意区分大小写。
//输入:
//"abccccdd"

//输出:
//7
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

//notice:创建一个map，用于存放各字符的个数。再遍历字符串，找到字符最大个偶数的值，再看有无奇数个落单，择情+1

func longestPalindrome(s string) int {
	if len(s) < 2 {
		return 1
	}
	countMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		countMap[string(s[i])] += + 1
	}
	fmt.Println(countMap)
	var countDoubleTimes int
	var countSingleTimes int
	if len(countMap) == 1 {
		for _, v := range countMap {
			return v
		}
	}
	for _, v := range countMap {
		if v%2 == 0 {
			countDoubleTimes += v
		} else {
			countDoubleTimes += v - 1
			countSingleTimes++
		}
	}
	if countSingleTimes > 0 {
		return countDoubleTimes + 1
	} else {
		return countDoubleTimes
	}
}

func main() {
	res := longestPalindrome("abb")
	fmt.Println(res)
}
