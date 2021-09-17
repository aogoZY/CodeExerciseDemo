package main

import (
	"fmt"
	"strings"
)

//翻转字符串里的单词
//输入："the sky is blue"
//输出："blue is sky the"

func main() {
	res := reverseWords("a good   example")
	fmt.Println(res)
}

//按空格切分原字符串为列表，对列表的元素从尾部开始遍历，若长度大于1则拼接到[]string的结果中，最后对[]string的结果用" "连起来返回。
func reverseWords(s string) string {
	splitRes := strings.Split(s, " ")
	var res []string
	for i := len(splitRes) - 1; i >= 0; i-- {
		if len(splitRes[i]) > 0 {
			res = append(res, splitRes[i])
		}
	}
	return strings.Join(res, " ")
}
