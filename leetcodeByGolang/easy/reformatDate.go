package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	num := strings.Split(date, " ")
	fmt.Println(num[0])
	fmt.Println(num[1])
	fmt.Println(num[2])
	MonthMap:=make(map[string]string)
	MonthMap["Jan"]="01"
	MonthMap["Feb"]="02"
	MonthMap["Mar"]="03"
	MonthMap["Apr"]="04"
	MonthMap["May"]="05"
	MonthMap["Jun"]="06"
	MonthMap["Jul"]="07"
	MonthMap["Aug"]="08"
	MonthMap["Sep"]="09"
	MonthMap["Oct"]="10"
	MonthMap["Nov"]="11"
	MonthMap["Dec"]="12"
	month :=MonthMap[num[1]]
	year :=num[2]
	length :=len(num[0])
	day :=num[0][:length-2]
	fmt.Println(day)
	if len(day)==1{
		day = "0"+day
	}
	res := year+"-"+month+"-"+day
	return res
}


func main() {
	res := reformatDate("6th Jun 1933")
	fmt.Println(res)   //2052-10-20"
}
