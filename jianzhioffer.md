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

2、其实是两个递增序列,若中间值>最右边的值,说明中间值在左边数组里,需要left左移

若中间值<最右边的值,说明右边的值不是最小值,left—1

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
   right := len(numbers)-1
   for left<=right{
      mid := (left + right) / 2
      if numbers[mid] > numbers[right]{
         left =mid +1
      }else {
         right = right -1
      }
   }
   return numbers[left]
}
```

