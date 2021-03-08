package main

import "fmt"

//[7,1,5,3,6,4]
// 7

//不限次数 啊哥！！！
func getMaxProfit(input []int) (res int) {
	var maxProfit int  //最大利润
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] > input[i] {
			maxProfit += input[i+1] - input[i]
		}
	}
	return maxProfit
}

func main() {
	input := []int{1,2,3,4,5}
	res := getMaxProfit2(input)
	fmt.Println(res)
}

func getMaxProfit2(input []int)(res int){
	var maxProfit int   //最大利润
	var smallValue int  //最小sell价格
	smallValue = input[0]
	for i:=0;i<len(input);i++{
		if input[i] < smallValue{
			smallValue = input[i]
		}
		if input[i]-smallValue>maxProfit{
			maxProfit=input[i]-smallValue
		}
	}
	res = maxProfit
	return res
}