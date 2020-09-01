# 剑指offer

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

