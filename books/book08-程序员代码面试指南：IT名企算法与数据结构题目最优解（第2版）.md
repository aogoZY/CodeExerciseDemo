### 程序员代码面试指南：IT名企算法与数据结构题目最优解（第2版）

左程云

# 链表

## 1、打印两个有序链表的公共部分

【题目】

给定两个有序链表的头指针head1和head2，打印两个链表的公共部分。

【解答】

> 从两个链表的头开始进行如下判断：
>
> · 如果head1的值小于head2，则head1往下移动。
>
> · 如果head2的值小于head1，则head2往下移动。
>
> · 如果head1的值与head2的值相等，则打印这个值，然后head1与head2都往下移动。
>
> · head1或head2有任何一个移动到null，则整个过程停止。

```
func PrintCommonPartNode(node1, node2 *Node) {
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			node1 = node1.Next
		} else if node1.Val > node2.Val {
			node2 = node2.Next
		} else {
			fmt.Println(node1.Val)
			node1 = node1.Next
			node2 = node2.Next
		}
	}
}
```



## 2、合并两个有序的单链表

【题目】

给定两个有序单链表的头节点 head1和 head2，请合并两个有序链表，合并后的链表依然有序，并返回合并后链表的头节点。

例如：

0->2->3->7->null

1->3->5->7->9->null

合并后的链表为：0->1->2->3->3->5->7->7->9->null

【解答】

假设两个链表的长度分别为M和N，直接给出时间复杂度为O(M+N)、额外空间复杂度为O(1)的方法。具体过程如下：

> 1．如果两个链表中有一个为空，说明无须合并过程，返回另一个链表的头节点即可。
>
> 2．比较head1和head2的值，小的节点是合并链表的头节点，记为head；在之后的步骤里，哪个链表的头节点的值更小，另一个链表的所有节点都会依次插入到这个链表中。
>
> 3．不妨设head节点所在的链表为链表1，另一个链表为链表2。链表1和链表2都从头部开始一起遍历，比较每次遍历到的两个节点的值，记为cur1和cur2，然后根据大小关系做出不同的调整，同时用一个变量pre表示上次比较时值较小的节点。
>
> 例如，链表1为1->5->6->null，链表2为2->3->7->null。
>
> cur1=1, cur2=2, pre=null。cur1小于cur2，不做调整，因为此时cur1较小，所以令pre=cur1=1
>
> cur1=5, cur2=2, pre=1。cur2小于cur1，让pre的next指针指向cur2, cur2的next指针指向cur1，这样，cur2便插入到链表1中。令pre=cur2=2，
>
> 链表1变为1->2->5->6->null，链表2变为3->7->null, cur1=5, cur2=3, pre=2。
>
> cur1=5, cur2=3, pre=2。此时又是cur2较小，链表1变为1->2->3->5->6->null，链表2为7->null, cur1=5, cur2=7, pre=3。
>
> cur1=5, cur2=7, pre=3。cur1小于cur2，pre=cur1=5。
>
> cur1=6, cur2=7, pre=5。cur1小于cur2，令pre=cur1=6，此时已经走到链表1的最后一个节点，再往下就结束，如果链表1或链表2有任何一个走到了结束，就进入步骤4。
>
> 4．如果链表1先走完，此时cur1=null, pre为链表1的最后一个节点，那么就把pre的next指针指向链表2当前的节点（即cur2），表示把链表2没遍历到的有序部分直接拼接到最后，调整结束。如果链表2先走完，说明链表2的所有节点都已经插入到链表1中，调整结束。
>
> 5．返回合并后链表的头节点head。
>
> 

```
package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

//给定两个有序单链表的头节点 head1和 head2，请合并两个有序链表，合并后的链表依然有序，并返回合并后链表的头节点。
//
//0->2->3->7->null
//
//1->3->5->7->9->null
//
//合并后的链表为：0->1->2->3->3->5->7->7->9->null

//notice:
//1．如果两个链表中有一个为空，说明无须合并过程，返回另一个链表的头节点即可。
//2．比较head1和head2的值，小的节点是合并链表的头节点，记为head；
// 在之后的步骤里，哪个链表的头节点的值更小，另一个链表的所有节点都会依次插入到这个链表中。
//3．不妨设head节点所在的链表为链表1，另一个链表为链表2。链表1和链表2都从头部开始一起遍历，比较每次遍历到的两个节点的值，
// 记为cur1和cur2，然后根据大小关系做出不同的调整，同时用一个变量pre表示上次比较时值较小的节点。
//4．如果链表1先走完，此时cur1=null, pre为链表1的最后一个节点，那么就把pre的next指针指向链表2当前的节点（即cur2），表示把链表2没遍历到的有序部分直接拼接到最后，
//  如果链表2先走完，说明链表2的所有节点都已经插入到链表1中，调整结束。
//5．返回合并后链表的头节点head。

func GetOrderNode(node1, node2 *Node) *Node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	var head, cur1, cur2 *Node
	// 比较两个链表的头节点 值更小的为head头
	//cur1、cur2分别为两个链表的当前值，我们默认头节点所在的链表为cur1

	if node1.Val < node2.Val {
		head = node1
		cur1 = node1
		cur2 = node2

	} else {
		head = node2
		cur1 = node2
		cur2 = node1
	}

	//pre节点存放每次cur1、cur2比较的最小值
	var pre, next *Node
	for cur1 != nil && cur2 != nil {
		//如果头节点所在链表的值小，需要把小值的位置存放到pre中，同时cur1链表往后移一位
		if cur1.Val < cur2.Val {
			pre = cur1
			cur1 = cur1.Next
		} else {
			//如果非头节点所在链表的值小，需要将非头节点接入头节点，小值存放到pre中，非头节点往后移一位
			next = cur2.Next
			pre.Next = cur2
			cur2.Next = cur1 //这两步是将当前节点接入头节点所在链表
			pre = cur2
			cur2 = next
		}
	}
	if cur1 == nil {
		pre.Next = cur2
	} else {
		pre.Next = cur1
	}
	return head
}

```



## 3、删除无序单链表中值重复出现的节点

【题目】

给定一个无序单链表的头节点head，删除其中值重复出现的节点。

例如：1->2->3->3->4->4->2->1->1->null，删除值重复的节点之后为1->2->3->4->null。

请按以下要求实现两种方法。

方法1：如果链表长度为N，时间复杂度达到O(N)。

方法2：额外空间复杂度为O(1)。

【解答】

方法一：利用哈希表。时间复杂度为O(N)，额外空间复杂度为O(N)。

具体过程如下：

1．生成一个哈希表，因为头节点是不用删除的节点，所以首先将头节点的值放入哈希表。

2．从头节点的下一个节点开始往后遍历节点，假设当前遍历到cur节点，先检查cur的值是否在哈希表中，如果在，则说明cur节点的值是之前出现过的，就将cur节点删除，删除的方式是将最近一个没有被删除的节点 pre 连接到cur的下一个节点，即pre.next=cur.next。如果不在，将cur节点的值加入哈希表，同时令pre=cur，即更新最近一个没有被删除的节点。

```
func removeRepNode(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	pre := node
	cur := node.Next
	nodeMap := make(map[int]bool)
	nodeMap[node.Val] = true
	for cur != nil {
		next := cur.Next
		_, ok := nodeMap[cur.Val]
		if ok {
			pre.Next = next
		} else {
			nodeMap[cur.Val] = true
			pre = cur
		}
		cur = next
	}
	return node
}
```

方法二：类似选择排序的过程，时间复杂度为O(N2)，额外空间复杂度为O(1)。

例如，链表1->2->3->3->4->4->2->1->1->null。

首先是头节点，节点值为1，往后检查所有值为1的节点，全部删除。链表变为：1->2->3->3->4->4->2->null。

然后是第二个节点，节点值为2，往后检查所有值为2的节点，全部删除。链表变为：1->2->3->3->4->4->null。

接着是第三个节点，节点值为3，往后检查所有值为3的节点，全部删除。链表变为：1->2->3->4->4->null。

最后是第四个节点，节点值为4，往后检查所有值为4的节点，全部删除。链表变为：1->2->3->4->null。

删除过程结束。

```
func removeRepNode2(node *Node) *Node {
	//cur 用于确定当前最外层遍历需要比较的节点，i
	cur := node
	//pre用于保存更新值最后的index位置
	//next用于循环内部的j
	var pre, next *Node
	for cur != nil {
		pre = cur
		next = cur.Next
		for next != nil {
			if cur.Val == next.Val {
				pre.Next = next.Next
			} else {
				pre = next
			}
			next = next.Next
		}
		cur = cur.Next
	}
	return node
}
```

4、两个单链表生成相加链表

public void removeRep2(Node head) {

Node cur = head;

Node pre = null;

Node next = null;

while (cur ! = null) {

pre = c

 2021-03-10 15:10:44

## 4、两个单链表生成相加链表

【题目】

```
输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
```

【解答】

将两个链表分别从左到右遍历，我们令 l1 和 l2 指向两个链表的头，用一个 tmp 值来存储同一位相加的结果，以及一个新的链表来存储 tmp 的值。此时没有指针指向生成的链表，所以需要将head先赋值为node，node需要将下一个节点的值node.next设为当前两链表数和，同时node需要向下移动指向当前节点。node.next=node

```
func addListReverse(node1, node2 *Node) *Node {
	res := &Node{0, nil}
	//	此处用head来指向生成链表的头节点
	head := res
	var temp int
	for node1 != nil || node2 != nil || temp > 0 {
		if node1 != nil {
			temp += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			temp += node2.Val
			node2 = node2.Next
		}
		// 链表的下一个节点为此节点两数和的余树
		res.Next = &Node{temp % 10, nil}
		//res节点后移
		res = res.Next
		//进位数
		temp = temp / 10
	}
	return head.Next
}
```



## 5、判断一个链表是否为回文结构

【题目】

给定一个链表的头节点head，请判断该链表是否为回文结构。

例如：

1->2->1，返回true。

1->2->2->1，返回true。

15->6->15，返回true。

1->2->3，返回false。

【解答】

用切片接受链表的值，对切片使用前后两个索引向中间夹，有不相等的值直接返回false

```
func IsPalindrome(node *Node) bool {
	var nodeList []int
	for node != nil {
		nodeList = append(nodeList, node.Val)
		node = node.Next
	}
	fmt.Println(nodeList)
	i := 0
	j := len(nodeList) - 1
	for i < j {
		if nodeList[i] != nodeList[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```











## 6、反转单向链表

【题目】

分别实现反转单向链表. 

输入：1->2->3 

输出：3->2->1 

【解答】

```
func reverseNode(node *Node) *Node {
	cur := node
	var pre *Node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		//pre, cur, cur.Next = cur, next, pre
	}
	return pre
}
```



## 7、在单链表删除倒数第K个节点

【题目】

分别实现两个函数，一个可以删除单链表中倒数第K个节点，另一个可以删除双链表中倒数第K个节点。

【解答】

使用双指针，快指针比慢指针速度快k个节点，当k到达链表尾部时慢指针所在位置即为倒数k的长度，若是快指针已经到链表尾k仍然大于零的话说明k本身大于链表长度,此时返回nil。

```
func removeLastKthNode(node *Node, k int) *Node {
	slow, fast := node, node
	for k > 0 {
		fast = fast.Next
		k--
		if fast == nil && k > 0 {
			return nil
		}
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
```



# 字符串

## 1、给定一个字符串str，判断是不是整体有效的括号字符串

【举例】

str="()"，返回true; str="(()())"，返回true; str="(())"，返回true。

str="())"。返回false; str="()("，返回false; str="()a()"，返回false。

补充问题：给定一个括号字符串str，返回最长的有效括号子串。

【举例】

str="(()())"，返回6; str="())"，返回2; str="()(()()("，返回4。

【解答】

法一：若是左括号就压入,若是右括号则和栈顶元素比较,相等的话pop栈顶元素,不等的话直接报错.

最终判断依据:stack是否为空.

```
//判断输入的字符是否可以相互抵消
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//用栈，判断当前值若为）是否和栈顶元素相等，相等则消消乐，不想等代表不合法直接return false
func isValid(s string) bool {
   var stack []byte
   paren_map := map[byte]byte{
      ')': '(',
      ']': '[',
      '}': '{',
   }
   if len(s) < 2 {
      return false
   }
   for i, ch := range s {
      if ch == '(' || ch == '[' || ch == '{' {
         stack = append(stack, s[i])
      } else if len(stack) == 0 || paren_map[s[i]] != stack[len(stack)-1] {
         return false
      } else {
         stack = stack[:len(stack)-1]
      }
   }
   if len(stack) == 0 {
      return true
   }
   return false
}

```

法二：相同的相消法

```
func isValid2(s string) bool {
   for {
      old := s
      s = strings.ReplaceAll(s, "()", "")
      s = strings.ReplaceAll(s, "[]", "")
      s = strings.ReplaceAll(s, "{}", "")
      if s== ""{
         return true
      }
      if len(s)==len(old){
         return false
      }
   }
   return false

}
```









## 2、数组中两个字符串的最小距离

【题目】

给定一个字符串数组strs，再给定两个字符串str1和str2，返回在strs中str1与str2的最小距离，如果str1或str2为null，或不在strs中，返回-1。

【举例】

strs=["1", "3", "3", "3", "2", "3", "1"], str1="1", str2="2"，返回2。

strs=["CD"], str1="CD", str2="AB"，返回-1。

进阶问题：如果查询发生的次数有很多，如何把每次查询的时间复杂度降为O(1)？

【解答】

原问题。从左到右遍历strs，用变量last1记录最近一次出现str1的位置，用变量last2记录最近一次出现str2的位置。如果遍历到str1，那么i-last2的值就是当前的str1和左边离它最近的str2之间的距离。如果遍历到str2，那么i-last1的值就是当前的str2和左边离它最近的str1之间的距离。用变量min记录这些距离的最小值即可。请参看如下的minDistance方法。

> public int minDistance(String[] strs, String str1, String str2) {
>
> if (str1 == null || str2 == null) {
>
> return -1;
>
> }
>
> if (str1.equals(str2)) {
>
> return 0;
>
> }
>
> int last1 = -1;
>
> int last2 = -1;
>
> int min = Integer.MAX_VALUE;
>
> for (int i = 0; i ! = strs.length; i++) {
>
> if (strs[i].equals(str1)) {
>
> min = Math.min(min, last2 == -1 ? min : i - last2);
>
> last1 = i;
>
> }
>
> if (strs[i].equals(str2)) {
>
> min = Math.min(min, last1 == -1 ? min : i - last1);
>
> last2 = i;
>
> }
>
> }
>
> return min == Integer.MAX_VALUE ? -1 : min;
>
> }
>
> 进阶问题。其实是通过数组strs先生成某种记录，在查询时通过记录进行查询。本书提供了一种记录的结构供读者参考，如果strs的长度为N，那么生成记录的时间复杂度为O(N2)，记录的空间复杂度为O(N2)，在生成记录之后，单次查询操作的时间复杂度可降为O(1)。本书实现的记录其实是一个哈希表HashMap<String, HashMap<String, Integer>>，这是一个key为string类型、value 为哈希表类型的哈希表。为了描述清楚，我们把这个哈希表叫作外哈希表，把 value代表的哈希表叫作内哈希表。外哈希表的key代表strs中的某种字符串，key所对应的内哈希表表示其他字符串到key字符串的最小距离。比如，当strs为["1", "3", "3", "3", "2", "3", "1"]时，生成的记录如下（外哈希表）：

如果生成了这种结构的记录，那么查询str1和str2的最小距离时只用两次哈希查询操作就可以完成。

如下代码的 Record 类就是这种记录结构的具体实现，建立记录过程就是 Record 类的构造函数，Record类中的minDistance方法就是做单次查询的方法。

public class Record {





## 3、翻转字符串

【题目】

给定一个字符类型的数组chas，请在单词间做逆序调整。只要做到单词的顺序逆序即可，对空格的位置没有特别要求。

【举例】

如果把chas看作字符串为"dog loves pig"，调整成"pig Loves dog"。

如果把chas看作字符串为"I'm a student."，调整成"student. a I'm"。

【解答】

- 按空格切分字符串为数组
- 从后往前遍历数组，长度大于0的元素拼接至[]string
- 整合[]string,用空格拼接成string返回

```
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
```



## [4、第一个只出现一次的字符](https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/)

【题目】

在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。【举例】

【事例】

```
s = "abaccdeff"
返回 "b"

s = "" 
返回 " "
```

【解答】

对字符串做第一遍遍历，使用长度为26的列表，index值对应当前字符串的字符，value对应当前字符串出现的总数

对字符串做第二遍遍历，若当前字符在列表中的value为1，代表仅出现一次，返回。

```
func firstUniqChar(s string) byte {
	list := make([]int, 26)
	for _, v := range s {
		list[v-'a'] += 1
	}
	fmt.Println(list)
	for _, value := range s {
		if list[value-'a'] == 1 {
			return byte(value)
		}
	}
	return ' '
}

```



# 树

```

```



# 其他

## [拼车](https://leetcode-cn.com/problems/car-pooling/)

这儿有一份乘客行程计划表 trips[][]，其中 trips[i] = [num_passengers, start_location, end_location] 包含了第 i 组乘客的行程信息：

必须接送的乘客数量；
乘客的上车地点；
以及乘客的下车地点。
这些给出的地点位置是从你的 初始 出发位置向前行驶到这些地点所需的距离（它们一定在你的行驶方向上）。请你根据给出的行程计划表和车子的座位数，来判断你的车是否可以顺利完成接送所有乘客的任务（当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false）。

```
输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false
```

```
输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true
```

分析：题目意思为在第一站有2个人上车，在第五站有五个人下车，在第三站有三个人上车，在第七站有三个人下车，问在整个过程中是否存在所有乘客树大于容量的情况。此时可以考虑用差分方式来做。差分即为后一个数与前一个数的差值，表示其变化。

[2, 0, 0, 0, -2]		对于第一列他的差分数据为

[0, 0, 3, 0,  0, 0, -3]	对于第二列差分数据为

[2, 0, 3, 0,  -2, 0, -3]		即总的差分数据为  所以计算出总的乘客数为[2,2,5,5,3,3,0]。所以容量最大为5，输入cap为5时返回true，cap为4时返回false。

```
func carPooling(trips [][]int, capacity int) bool {
	res := make([]int, 1024)
	var length int
	for j := 0; j < len(trips); j++ {
		pass := trips[j][0]
		start := trips[j][1]
		end := trips[j][2]
		res[start-1] += pass
		res[end-1] -= pass
		if end > length {
			length = end
		}
	}
	countRes := make([]int, length)
	countRes[0] = res[0]
	for i := 1; i < length; i++ {
		countRes[i] = countRes[i-1] + res[i]
		if countRes[i] > capacity {
			return false
		}
	}
	return true
}
```

## [航班预订统计](https://leetcode-cn.com/problems/corporate-flight-bookings/)

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。

> 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
> 输出：[10,55,45,25,25]
> 解释：
> 航班编号        1   2   3   4   5
> 预订记录 1    10  10
> 预订记录 2 ：       20  20
> 预订记录 3：        25  25  25  25
> 总座位数：     10  55  45  25  25
> 因此，answer = [10,55,45,25,25]



暴力解法

```
func corpFlightBookings(bookings [][]int, n int) []int {
	length := len(bookings)
	countMap := make(map[int]int)
	for j := 0; j < length; j++ {
		start := bookings[j][0]
		end := bookings[j][1]
		for start <= end {
			countMap[start] += bookings[j][2]
			start++
		}
	}
	var res []int
	for i := 1; i <= n; i++ {
		res = append(res, countMap[i])
	}
	return res
}
```

差分解法

```
//差分管理
func corpFlightBookings2(bookings [][]int, n int) []int {
	length := len(bookings)
	res := make([]int, n)
	for j := 0; j < length; j++ {
		start := bookings[j][0]
		end := bookings[j][1]
		for i := start-1; i < end; i++ {
			res[i] += bookings[j][2]
			start++
		}
	}
	return res
}

```



start

年轻人总会找借口说这个东西不是我感兴趣的，所以做不好是应该的。但他们没有注意的是，你面对的事情中感兴趣的事情总是少数，这就使得大多数时候你做事情的态度总是很懈怠、很消极，这使你变成了一个懈怠的人。当你真正面对自己感兴趣的东西时，你发现你已经攥不紧拳头了

