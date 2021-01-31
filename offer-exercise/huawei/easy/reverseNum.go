package main

import (
	"fmt"
)

func ResevieNum(num int){
	num_str:=fmt.Sprintf("%d",num)
	for i:=len(num_str)-1;i>=0;i--{
		fmt.Print(string(num_str[i]))
	}
	fmt.Println()

}

func main(){
	var num int
	//fmt.Scanln(&num)
	num=1516000
	ResevieNum(num)

}