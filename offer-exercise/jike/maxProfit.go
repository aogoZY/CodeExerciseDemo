package main

import "fmt"

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

//思路
//1、DFS O(n*2)
//2、贪心算法 O(n)一遍遍历，下一次比这一次的多就买入
//3、DP O(n)
func main() {
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res)
}

func maxProfit(input []int) int {
	small:=input[0]
	maxProfit:=0
	for i:=0;i<len(input);i++{
		if input[i]<small{
			small = input[i]
		}
		nowProfit:=input[i]-small
		if nowProfit>maxProfit{
			maxProfit=nowProfit
		}
	}
	return maxProfit
}
