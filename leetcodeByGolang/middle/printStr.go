package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "a(b(c)<2>(d)<3>)<2>e"
	//res,index:=GetContentByCharacter(str,"(",")",0)
	for checkHasKuohao(str) {
		str = printStr(str)
	}
	fmt.Println(str)
	//fmt.Println(index)
	//fmt.Println(res)
	//res2,index2:=GetContentByCharacter(str,"<",">",index+1)
	//fmt.Println(res2)
	//fmt.Println(index2)

}

func checkHasKuohao(str string) (res bool) {
	if strings.Contains(str, "(") {
		return true
	}
	return false
}

func printStr(str string) (res string) {
	content, lastIndex := GetContentByCharacter(str, "(", ")", 0)
	repeatNum, _ := GetContentByCharacter(str, "<", ">", lastIndex+1)
	num, _ := strconv.Atoi(repeatNum)
	contentNote := GetContentNotes(content, num)
	replaceContent := "(" + content + ")<" + repeatNum + ">"
	res = strings.ReplaceAll(str,  replaceContent,contentNote)
	return res
}

func GetContentNotes(content string, num int) (res string) {
	for i := 0; i < num; i++ {
		res = res + content
	}
	return res
}

func GetContentByCharacter(str, left, right string, index int) (content string, lastIndex int) {
	var isFirst bool = true
	var leftIndex, rightIndex int
	var countNum int
	for i := index ; i <= len(str); i++ {
		if string(str[i]) == left {
			if isFirst {
				leftIndex = i
				isFirst = false
			}
			countNum++
		} else if string(str[i]) == right {
			countNum--
			if countNum == 0 {
				rightIndex = i
				break
			}
		}
	}
	content = str[leftIndex+1:rightIndex]
	return content, rightIndex
}
