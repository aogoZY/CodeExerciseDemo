package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("pwwkew")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	var (
		length int //最后返回值
		begin  int //每个子串的起始位置
	)
	//使用这个map来判定当前子串中是否出现了重复的字符
	tmpMap := make(map[string]int, len(s))
	for i := 0; i < len(s); i++ { //只遍历这一次，没有其他循环，不然就超出时间限制
		//这一步很关键，首先使用OK来判断当前子串中是否存在循环当前遍历到的字符，如果不存在，就把这个字符和它的位置加入map
		//如果这个字符存在于map中，那么就判断这个字符的位置和begin谁更靠前
		//如果这个字符的位置更靠前，说明这个字符出现的位置不是当前子串，那么也相当于在当前子串中第一次出现，然后更新map中这个字符的位置
		//如果begin更靠前，说明在当前子串中出现了重复字符，需要计算当前无重复字符子串的长度了
		if index, OK := tmpMap[string(s[i])]; !OK || index < begin {
			tmpMap[string(s[i])] = i //更新位置
			if i == len(s)-1 { //这里计算的长度是给定字符串的最后一个子串的长度
				if (i - begin + 1) > length {
					length = (i - begin + 1)
				}
			}
		} else { //当遇到重复字符的时候，计算子串的长度
			if (i - begin) > length {
				length = (i - begin)
			}
			//更新下一个子串的起始位置
			begin = tmpMap[string(s[i])] + 1
			//更新遍历到的字符在map中的位置
			tmpMap[string(s[i])] = i
		}
	}
	return length
}
