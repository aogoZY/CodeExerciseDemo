# 剑指offer

[TOC]



# easy

#### 剑指 Offer 03. 数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

> 示例 1：
>
> 输入：
> [2, 3, 1, 0, 2, 5, 3]
> 输出：2 或 3 

```go
func findRepeatNumber(nums []int) int {
	for i,item :=range nums{
		for i:=i+1;i<len(nums);i++{
			if nums[i]==item{
				return item
			}
		}
	}
	return 0
}
```

#### [剑指 Offer 04. 二维数组中的查找]

在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

> 示例:
>
> 现有矩阵 matrix 如下：
>
> [
>   [1,   4,  7, 11, 15],
>   [2,   5,  8, 12, 19],
>   [3,   6,  9, 16, 22],
>   [10, 13, 14, 17, 24],
>   [18, 21, 23, 26, 30]
> ]
> 给定 target = 5，返回 true。
>
> 给定 target = 20，返回 false。
>

```go
func findNumberIn2DArray(arr [][]int, target int) bool {
	if len(arr)!=0 && len(arr[0])!=0{
		width := len(arr)
		length := len(arr[0])
		i := 0
		j := length - 1
		for i <= width-1 && j >= 0 {
			if target < arr[i][j] {
				j--
			} else if target > arr[i][j] {
				i++
			} else if target == arr[i][j] {
				return true
			}
		}
	}
	return false
}
```

#### [剑指 Offer 05. 替换空格]

**Notice:字符串可直接遍历,字符串可直接拼接str+=“%20”**

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

> 示例 1：
>
> 输入：s = "We are happy."
> 输出："We%20are%20happy."

```go
func replaceSpace(s string) string {
	var result string
	for _, item := range s {
		if string(item) == " " {
			result += "%20"
		} else {
			result += string(item)
		}
	}
	fmt.Println(result)
	return result
}
```

#### [剑指 Offer 10- I. 斐波那契数列](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/)

**Notice:**

- **直接递归会超过时间复杂度.可考虑for遍历将每个值存到list里.**

- **为何需要对 1e9+7取模:题目的答案是很大的，要让答案落在整型的范围内----对一个很大的质数取模即可。对质数取模的话，能尽可能地避免模数相同的数之间具备公因数，来达到减少冲突的目的。**

> 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：
>
> F(0) = 0,   F(1) = 1
> F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
> 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
>
> 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
>

```go
func fib(n int) int {
	var res map[int]int
	res = make(map[int]int)
	res[0] = 0
	res[1] = 1
	if n >= 2 {
		for i := 2; i <=n; i++ {
			res[i] = (res[i-1] + res[i-2]) %  1000000007
		}
		return res[n]
	}
	return res[n]
}
```

#### [剑指 Offer 10- II. 青蛙跳台阶问题](https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/)

> 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
>
> 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
>
> 示例 1：
>
> 输入：n = 2
> 输出：2
> 示例 2：
>
> 输入：n = 7
> 输出：21

``` go
func numWays(n int) int {
   if n== 1 || n==2{
		return n
	}
	if n==0{
		return 1
	}
	var res map[int]int
	res =make(map[int]int)
	for i:=3;i<=n;i++{
		res[1]=1
		res[2]=2
		res[i] = (res[i-2]+res[i-1]) % 1000000007
	}
	return res[n]

```

#### [剑指 Offer 11. 旋转数组的最小数字](https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/)

Notice:

1、暴力解法

2、可以看成是两个递增序列,找到两者之间的分段点,分段点后的第一位数则为最小数.

采用二分法:

若中间值>最右边的值,说明分段点在后半截,区间取后半段,left = mid +1;

若中间值<最右边的值,说明分段点在前半截,区间取前半段,right = mid;

若中间值和右边值相等,则left = left -1,再做判断

> 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  
>
> 示例 1：
>
> 输入：[3,4,5,1,2]
> 输出：1
> 示例 2：
>
> 输入：[2,2,2,0,1]
> 输出：0

```go
func minArray2(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] == numbers[right] {
			right = right - 1

		}
	}
	return numbers[left]
}
```

#### [剑指 Offer 15. 二进制中1的个数](https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/)

请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。

每次将num位移,比较其和1的关系

> 示例 1：
>
> 输入：00000000000000000000000000001011
> 输出：3
> 解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
> 示例 2：
>
> 输入：00000000000000000000000010000000
> 输出：1
> 解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
> 示例 3：
>
> 输入：11111111111111111111111111111101
> 输出：31
> 解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

```
func hammingWeight(num uint32) int {
	var res int
	for num >0{
		if num & 1 == 1{
			res++
		}
		num = num >>1
	}
	return res
}
```

#### [剑指 Offer 16. 数值的整数次方](https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)

实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

> 示例 1:
>
> 输入: 2.00000, 10
> 输出: 1024.00000
> 示例 2:
>
> 输入: 2.10000, 3
> 输出: 9.26100
> 示例 3:
>
> 输入: 2.00000, -2
> 输出: 0.25000
> 解释: 2-2 = 1/22 = 1/4 = 0.25

#### [剑指 Offer 17. 打印从1到最大的n位数](https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/)

难度简单47收藏分享切换为英文关注反馈

输入数字 `n`，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

> 输入: n = 1
> 输出: [1,2,3,4,5,6,7,8,9]

```go
func printNumbers(n int) []int {
   var big string
   var res []int
   for i := 0; i < n; i++ {
      big = big + "9"
   }
   fmt.Println(big)
   intNumI,_ :=strconv.Atoi(big)
   for i:=1;i<=intNumI;i++{
      res= append(res, i)
   }
   return res
}
```

#### [1507. 转变日期格式](https://leetcode-cn.com/problems/reformat-date/)

给你一个字符串 `date` ，它的格式为 `Day Month Year` ，其中：

- `Day` 是集合 `{"1st", "2nd", "3rd", "4th", ..., "30th", "31st"}` 中的一个元素。
- `Month` 是集合 `{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}` 中的一个元素。
- `Year` 的范围在 `[1900, 2100]` 之间。

请你将字符串转变为 `YYYY-MM-DD` 的格式，其中：

- `YYYY` 表示 4 位的年份。

- `MM` 表示 2 位的月份。

- `DD` 表示 2 位的天数。

  > 输入：date = "20th Oct 2052"
  > 输出："2052-10-20"
  > 示例 2：
  >
  > 输入：date = "6th Jun 1933"
  > 输出："1933-06-06"
  > 示例 3：
  >
  > 输入：date = "26th May 1960"
  > 输出："1960-05-26"



#### [455. 分发饼干](https://leetcode-cn.com/problems/assign-cookies/)

假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

**注意：**

你可以假设胃口值为正。
一个小朋友最多只能拥有一块饼干。

> 输入: [1,2,3], [1,1]
>
> 输出: 1
>
> 解释: 
> 你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
> 虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
> 所以你应该输出1。

> 输入: [1,2], [1,2,3]
>
> 输出: 2
>
> 解释: 
> 你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
> 你拥有的饼干数量和尺寸都足以让所有孩子满足。
> 所以你应该输出2.

```go
func findContentChildren(g []int, s []int) int {
	sortG := sort(g)
	fmt.Println(sortG)
	sortS := sort(s)
	fmt.Println(sortS)
	var res int
	var i, j int
	for i <= len(sortG) -1 && j <= len(sortS)-1 {
		if sortG[i] <= sortS[j] {
			i++
			j++
			res++
		} else {
			j++
		}
	}
	return res
}

func sort(raw []int) (sortedList []int) {
	for i := len(raw) - 1; i > 0; i-- {
		for j := 0; j < len(raw)-1; j++ {
			if raw[j] > raw[j+1] {
				b := raw[j]
				raw[j] = raw[j+1]
				raw[j+1] = b
			}
		}
	}
	return raw
}
```



#### [258. 各位相加](https://leetcode-cn.com/problems/add-digits/)

难度简单275收藏分享切换为英文关注反馈

给定一个非负整数 `num`，反复将各个位上的数字相加，直到结果为一位数。

**示例:**

> 输入: 38
> 输出: 2 
> 解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。

```go
func addDigits(num int) int {
   if num < 10 {
      return num
   }
   return addDigits(GetSum(num))
}

func GetSum(num int) int {
   var rest int = num
   var count int
   for rest > 0 {
      weishu := rest % 10
      rest = rest / 10
      count += weishu
   }
   return count
}
```





# Middle

#### [1508. 子数组和排序后的区间和](https://leetcode-cn.com/problems/range-sum-of-sorted-subarray-sums/)

给你一个数组 nums ，它包含 n 个正整数。你需要计算所有非空连续子数组的和，并将它们按升序排序，得到一个新的包含 n * (n + 1) / 2 个数字的数组。

请你返回在新数组中下标为 left 到 right （下标从 1 开始）的所有数字和（包括左右端点）。由于答案可能很大，请你将它对 10^9 + 7 取模后返回。

>  示例 1：
>
> 输入：nums = [1,2,3,4], n = 4, left = 1, right = 5
> 输出：13 
> 解释：所有的子数组和为 1, 3, 6, 10, 2, 5, 9, 3, 7, 4 。将它们升序排序后，我们得到新的数组 [1, 2, 3, 3, 4, 5, 6, 7, 9, 10] 。下标从 le = 1 到 ri = 5 的和为 1 + 2 + 3 + 3 + 4 = 13 。
> 示例 2：
>
> 输入：nums = [1,2,3,4], n = 4, left = 3, right = 4
> 输出：6
> 解释：给定数组与示例 1 一样，所以新数组为 [1, 2, 3, 3, 4, 5, 6, 7, 9, 10] 。下标从 le = 3 到 ri = 4 的和为 3 + 3 = 6 。
>

#### [260. 只出现一次的数字 III](https://leetcode-cn.com/problems/single-number-iii/)

难度中等280收藏分享切换为英文关注反馈

给定一个整数数组 `nums`，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

**示例 :**

```
输入: [1,2,1,3,2,5]
输出: [3,5]
```

**注意：**

1. 结果输出的顺序并不重要，对于上面的例子， `[5, 3]` 也是正确答案。
2. 你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

```go
func singleNumber(nums []int) []int {
   var res []int
   var testMap map[int]int
   testMap = make(map[int]int)
   for _, item := range nums {
      _, ok := testMap[item]
      if ok {
         testMap[item] += 1
      } else {
         testMap[item] = 1
      }

   }
   fmt.Println(testMap)
   for item := range testMap {
      if testMap[item]==1{
         res= append(res, item)
      }
   }
   return res
}
```

#### 输入根据数字输出重复字符串

输入a(b(c)<2>(d)<3>)<2>e,输出abccdddbccddde的字符串

1. 找到需要重复的内容

2. 获取需要重复的次数

3. 拼接重复的内容,拼成新的字符串,replace掉旧的内容

4. 判断是否还存在(,若有则一直按上述流程执行.

   

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "a(b(c)<2>(d)<3>)<2>e"
	for checkHasKuohao(str) {
		str = printStr(str)
	}
	fmt.Println(str)
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

```

#### 查找查找两个字符串a,b中的最长公共子串。

若有多个，输出在较短串中最先出现的那个。

```go
package main

import (
   "fmt"
   "strings"
)

//思路：先遍历最短的子串，看子串是否在长子串中
func main() {
   strA := "abcdefghijklmnop"
   strB := "bcsafjklmnopqrstuvw"
   res := printFirstLongSubstr(strA, strB)
   fmt.Println(res)
}

func printFirstLongSubstr(strA, strB string) (res string) {
   if len(strA) > len(strB) {
      strA, strB = strB, strA
   }
   length := 0
   for i := 0; i < len(strA); i++ {
      for j := i + 1; j < len(strA); j++ {
         subStr := strA[i:j]
         if strings.Contains(strB, subStr) && (j-i) > length {
            res = subStr
            length = j - i
         }
      }
   }
   return res
}
```

#### 取出字符串中连续最长的数字子串

输出字符串中最长的数字字符串和它的长度，中间用逗号间隔。如果有相同长度的串，则要一块儿输出（中间不间隔），但是长度还是一串的长度，与数字字符串间用逗号间隔

```
input:abcd12345ed125ss123058789
output:123058789,9
```



```go
package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
   "strings"
)

//输出字符串中最长的数字字符串和它的长度，中间用逗号间隔。
// 如果有相同长度的串，则要一块儿输出（中间不间隔），但是长度还是一串的长度，与数字字符串间用逗号间隔。

//abcd12345ed125ss123058789
//123058789,9

//思路1：遍历字符串，对于数字类型用temp变量记录长度，取当前str【i-temp+1：i】为数字字符子串。
//若下一个子数字子串出现，比较其与当前maxLength长度，长则置换。相等取其res+=str【i-temp+1：i】

//思路2：将不是数字的值全部换成'A'，再用'A'split字符串。遍历数组，拿到最大长度的值，在遍历一次，将长度等于最大值的str取出来

func main() {
   sc := bufio.NewScanner(os.Stdin)
   for sc.Scan() {
      str := sc.Text()
      res, max := GetLongNumAndLength2(str)
      fmt.Println(res + "," + strconv.Itoa(max))
   }
   //var str string
   //str = "ad12345ed125ss123058789"
   //fmt.Scan(&str)
   //res, length := GetLongNumAndLength2(str)
   //lengthStr := strconv.Itoa(length)
   //outPut := res + "," + lengthStr
   //fmt.Print(outPut)
}

func GetLongNumAndLength(str string) (res string, length int) {
   if len(str) < 1 {
      return
   }
   var temp int
   var max int
   for i := 0; i < len(str); i++ {
      //fmt.Println(str[i])
      if str[i] < '0' || str[i] > '9' {
         temp = 0
      } else {
         temp++
         if temp > max {
            max = temp
            res = str[i-temp+1 : i+1]
         } else if temp == max {
            res += str[i-temp+1 : i+1]
         }
      }
   }
   return res, max
}

func GetLongNumAndLength2(str string) (res string, length int) {
   var fmtStr string
   for i := 0; i < len(str); i++ {
      if str[i] < '0' || str[i] > '9' {
         fmtStr += "A"
      } else {
         fmtStr += string(str[i])
      }
   }
   //fmt.Println(fmtStr)
   splitRes := strings.Split(fmtStr, "A")
   max := 0
   for i := 0; i < len(splitRes); i++ {
      if max < len(splitRes[i]) {
         max = len(splitRes[i])
      }
   }
   for i := 0; i < len(splitRes); i++ {
      if len(splitRes[i]) == max {
         res += splitRes[i]
      }
   }

   return res, max
}
```

#### 给定一个数组，求股票买卖最佳值

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。

```go
输入: [7,1,5,3,6,4]
输出: 5
```

```go
输入: [7,6,4,3,1]
输出: 0
```

```go
package main

import (
   "fmt"
)

//买卖股票 只允许操作1次
//思路一：暴力解法 两个for 循环记录每两个值之间的差价 mark差值最大的为profit
//思路二：将股票价当成一个折线图,只遍历一次，其实找到最低点和最高的差值就可以了。遍历过程中寻找最小的值，并求任意值和最小值之间的差的最大值。
func main() {
   input := []int{7, 1, 5, 3, 6, 4, 1}
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
```

# Tree

#### [ 二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

**说明：**叶子节点是指没有子节点的节点。

![img](https://assets.leetcode.com/uploads/2020/10/12/ex_depth.jpg)

输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

notice:递归求其左子树、右子树的最小深度+1

```
package main

import "fmt"

type TreeNode struct {
   Val   int
   Left  *TreeNode
   Right *TreeNode
}
//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//notice:可理解为：若求f（n）则求f（n-1），当前节点的左、右子树的最小值+1
func minDepth(root *TreeNode) int {
   if root == nil {
      return 0
   }
   dl := minDepth(root.Left)
   dr := minDepth(root.Right)
   if root.Left == nil {
      return dr + 1
   } else if root.Right == nil {
      return dl + 1
   }else {
      return min(dl,dr)+1
   }
}

func min(x, y int) int {
   if x > y {
      return y
   }
   return x
}
```

#### [112. 路径总和](https://leetcode-cn.com/problems/path-sum/)

给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

![image-20201104234239293](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201104234239293.png)

Notice:反向倒推:若是左子树和为sum-当前value 或者右子树和为sum-当前value都可.

询问是否存在从当前节点 root 到叶子节点的路径，满足其路径和为 sum。假定从根节点到当前节点的值之和为 val，我们可以将这个大问题转化为一个小问题：是否存在从当前节点的子节点到叶子的路径，满足其路径和为 sum - val。满足递归的性质，若当前节点就是叶子节点，那么我们直接判断 sum 是否等于 val 即可（因为路径和已经确定，就是当前节点的值，我们只需要判断该路径和是否满足条件）。若当前节点不是叶子节点，我们只需要递归地询问它的子节点是否能满足条件即可

```go
func hasPathSum(root *TreeNode, sum int) bool {
   if root == nil {
      return false
   }
   if root.Left == nil && root.Right == nil {
      return sum == root.Val
   }
   return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
```