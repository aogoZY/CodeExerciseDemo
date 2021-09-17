package string

import (
	"fmt"
	sort2 "sort"
)

//买卖股票 只允许操作1次
//思路一：暴力解法 两个for 循环记录每两个值之间的差价 mark差值最大的为profit
//思路二：将股票价当成一个折线图,只遍历一次，其实找到最低点和最高的差值就可以了。遍历过程中寻找最小的值，并求任意值和最小值之间的差的最大值。
func main() {
	input := []int{1,2,4,2,5,7,2,4,9,0}
	res := maxProfit2(input)
	fmt.Println(res)
}

func maxProfit(prices []int) (res int) {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > res {
				res = prices[j] - prices[i]
			}
		}
	}
	return res
}

func maxProfit2(prices []int) (res int) {
	if len(prices) == 0 || prices == nil {
		return 0
	}
	small := prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i]-small > max {
			max = prices[i] - small
		}
		if prices[i] < small {
			small = prices[i]
		}
	}
	return max
}

//买卖股票 允许操作多次
//思路：仍然看成一个折线图 将其连续上升的多个子区间的最大差值找出来即可
//判断逻辑是：只要判断下一天的价格比今天的高 我就今天买明天卖
func maxProfit3(prices []int) (res int) {
	if len(prices) == 0 {
		return 0
	}
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

//你最多可以完成 两笔 交易。
//[3,3,5,0,0,3,1,4]
//思路：找出子串的最大差值 存起来 再去其中前两位差值求和
//1,2,4,2,5,7,2,4,9,0
func maxProfit4(prices []int) (res int) {
	var num int
	var storeList []int
	for i := 0; i < len(prices)-1; i++ {
		if i < len(prices)-2 && prices[i+1] > prices[i] {
			num += prices[i+1] - prices[i]
		} else if i == len(prices)-2 && (prices[i+1] > prices[i]) {
			num += prices[i+1] - prices[i]
			storeList = append(storeList, num)
		} else {
			storeList = append(storeList, num)
			num = 0
		}

	}
	//fmt.Println(storeList)
	if len(storeList) == 0 || storeList == nil {
		return 0
	}
	sort2.Ints(storeList)
	//fmt.Println(storeList)
	if len(storeList) == 1 {
		return storeList[0]
	}
	return storeList[len(storeList)-1] + storeList[len(storeList)-2]
}
