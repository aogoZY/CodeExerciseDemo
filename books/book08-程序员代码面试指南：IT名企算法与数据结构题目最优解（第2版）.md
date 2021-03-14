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

找到字符串的最长无重复字符子串

【题目】

给定一个字符串str，返回str的最长无重复字符子串的长度。

【举例】

str="abcd"，返回4。

str="aabcb"，最长无重复字符子串为"abc"，返回3。

【要求】

如果str的长度为N，请实现时间复杂度为O(N)的方法。

【解答】

如果str长度为N，字符编码范围是M，本题可做到时间复杂度为O(N)，额外空间复杂度为O(M)。下面介绍这种方法的具体实现。

1．在遍历str之前，先申请几个变量。哈希表map, key表示某个字符，value为这个字符最近一次出现的位置。整型变量pre，如果当前遍历到字符str[i], pre表示在必须以str[i-1]字符结尾的情况下，最长无重复字符子串开始位置的前一个位置，初始时pre=-1。整型变量len，记录以每一个字符结尾的情况下，最长无重复字符子串长度的最大值，初始时，len=0。从左到右依次遍历 str，假设现在遍历到 str[i]，接下来求在必须以 str[i]结尾的情况下，最长无重复字符子串的长度。

2.map(str[i])的值表示之前的遍历中最近一次出现str[i]字符的位置，假设在a位置。想要求以str[i]结尾的最长无重复子串，a位置必然不能包含进来，因为str[a]等于str[i]。

3．根据pre的定义，pre+1表示在必须以str[i-1]字符结尾的情况下，最长无重复字符子串的开始位置。也就是说，以str[i-1]结尾的最长无重复子串是向左扩到pre位置停止的。

4．如果pre位置在a位置的左边，因为str[a]不能包含进来，而str[a+1..i-1]上都是不重复的，所以以str[i]结尾的最长无重复字符子串就是str[a+1..i]。如果pre位置在a位置的右边，以str[i-1]结尾的最长无重复子串是向左扩到 pre 位置停止的。所以以 str[i]结尾的最长无重复子串向左扩到pre位置也必然会停止，而且str[pre+1..i-1]这一段上肯定不含有str[i]，所以以str[i]结尾的最长无重复字符子串就是str[pre+1..i]。

5．计算完长度之后，pre位置和a位置哪一个在右边，就作为新的pre值。然后计算下一个位置的字符，整个过程中求得所有长度的最大值用len记录下来返回即可。

具体请参看如下代码中的maxUnique方法。

> public int maxUnique(String str) {
>
> if (str == null || str.equals("")) {
>
> return 0;
>
> }
>
> char[] chas = str.toCharArray();
>
> int[] map = new int[256];
>
> for (int i = 0; i < 256; i++) {
>
> map[i] = -1;
>
> }
>
> int len = 0;
>
> int pre = -1;
>
> int cur = 0;
>
> for (int i = 0; i ! = chas.length; i++) {
>
> pre = Math.max(pre, map[chas[i]]);
>
> cur = i - pre;
>
> len = Math.max(len, cur);
>
> map[chas[i]] = i;
>
> }
>
> return len;
>
> }
>
>  2021-03-09 16:19:43

> solution
>
> public class MyComparator implements Comparator<String> {
>
> @Override
>
> public int compare(String a, String b) {
>
> return (a + b).compareTo(b + a);
>
> }
>
> }
>
> public String lowestString(String[] strs) {
>
> if (strs == null || strs.length == 0) {
>
> return "";
>
> }
>
> // 根据新的比较方式排序
>
> Arrays.sort(strs, new MyComparator());
>
> String res = "";
>
> for (int i = 0; i < strs.length; i++) {
>
> res += strs[i];
>
> }
>
> return res;
>
> }



题目

拼接所有字符串产生字典顺序最小的大写字符串

【题目】

给定一个字符串类型的数组 strs，请找到一种拼接顺序，使得将所有的字符串拼接起来组成的大写字符串是所有可能性中字典顺序最小的，并返回这个大写字符串。

【举例】

strs=[ "abc", "de" ]，可以拼成"abcde"，也可以拼成"deabc"，但前者的字典顺序更小，所以返回"abcde"。

strs=["b", "ba" ]，可以拼成"bba"，也可以拼成"bab"，但后者的字典顺序更小，所以返回"bab"。



【解答】

有一种思路为：先把strs中的字符串按照

 

括号字符串的有效性和最长有效长度

# 

## 给定一个字符串str，判断是不是整体有效的括号字符串

【举例】

str="()"，返回true; str="(()())"，返回true; str="(())"，返回true。

str="())"。返回false; str="()("，返回false; str="()a()"，返回false。

补充问题：给定一个括号字符串str，返回最长的有效括号子串。

【举例】

str="(()())"，返回6; str="())"，返回2; str="()(()()("，返回4。

【难度】

原问题 士 ★☆☆☆

补充问题 尉 ★★☆☆

【解答】

原问题。判断过程如下：

1．从左到右遍历字符串str，判断每一个字符是不是’(’或’)'，如果不是，就直接返回false。

2．遍历到每一个字符时，都检查到目前为止’(’和’)’的数量，如果’)’更多，则直接返回false。

3．遍历后检查’(’和’)’的数量，如果一样多，则返回true，否则返回false。

具体过程参看如下代码中的isValid方法。

> public boolean isValid(String str) {
>
> if (str == null || str.equals("")) {
>
> return false;
>
> }
>
> char[] chas = str.toCharArray();
>
> int status = 0;
>
> for (int i = 0; i < chas.length; i++) {
>
> if (chas[i] ! = ')' && chas[i] ! = '(') {
>
> return false;
>
> }
>
> if (chas[i] == ')' && --status < 0) {
>
> return false;
>
> }
>
> if (chas[i] == '(') {
>
> status++;
>
> }
>
> }
>
> return status == 0;
>
> }
>
> 补充问题。用动态规划求解，可以做到时间复杂度为O(N)，额外空间复杂度为O(N)。首先生成长度和str字符串一样的数组dp[], dp[i]值的含义为str[0..i]中必须以字符str[i]结尾的最长的有效括号子串长度。那么dp[i]值可以按如下方式求解：

1.dp[0]=0。只含有一个字符肯定不是有效括号字符串，长度自然是0。

2．从左到右依次遍历str[1..N-1]的每个字符，假设遍历到str[i]。

3．如果str[i]=='('，有效括号字符串必然是以’)’结尾，而不是以’(’结尾，所以dp[i] = 0。

4．如果str[i]==')'，那么以str[i]结尾的最长有效括号子串可能存在。dp[i-1]的值代表必须以str[i-1]结尾的最长有效括号子串的长度，所以，如果i-dp[i-1]-1位置上的字符是’('，就能与当前位置的str[i]字符再配出一对有效括号。比如"(()())"，假设遍历到最后一个字符’)'，必须以倒数第二个字符结尾的最长有效括号子串是"()()"，找到这个子串之前的字符，即i-dp[i-1]-1位置的字符，发现是’('，所以它可以和最后一个字符再配出一对有效括号。如果该情况发生，dp[i]的值起码是dp[i-1]+2，但还有一部分长度容易被人忽略。比如，"()(())"，假设遍历到最后一个字符’)'，通过上面的过程找到的必须以最后字符结尾的最长有效括号子串起码是"(())"，但是前面还有一段"()"，可以和"(())"结合在一起构成更大的有效括号子串。也就是说，str[i-dp[i-1]-1]和str[i]配成了一对，这时还应该把dp[i-dp[i-1]-2]的值加到dp[i]中，这么做表示把str[i-dp[i-1]-2]结尾的最长



数组中两个字符串的最小距离

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



## 删除多余字符得到字典序最小的字符串

【题目】

给定一个全是小写字母的字符串str，删除多余字符，使得每种字符只保留一个，并让最终结果字符串的字典序最小。

【举例】

str = "acbc"，删掉第一个’c'，得到"abc"，是所有结果字符串中字典序最小的。

str = "dbcacbca"，删掉第一个’b'、第一个’c'、第二个’c'、第二个’a'，得到"dabc"，是所有结果字符串中字典序最小的。

【解答】

不考虑怎么去删除，应考虑怎么去挑选。str的结果字符串记为res，假设str长度为N，其中有 K 种不同的字符，那么 res 长度为 K。思路是怎么在 str 中从左到右依次挑选出 res[0]、res[1]、...、res[K-1]。举个例子，str[0..9]="baacbaccac"，一共3种字符，所以要在str中从左到右依次找到res[0..2]。

1．建立str[0..9]的字频统计，b有2个、a有4个、c有4个。

2．从左往右遍历 str[0..9]，遍历到字符的字频统计减1，当发现某一种字符的字频统计已经为0时，遍历停止。在例子中当遍历完"baacb"时，字频统计为b有0个、a有2个、c有3个，发现b的字频已经为0，所以停止遍历，当前遍历到str[4]。str[5..9]为"accac"已经没有b了，而流程是在str中从左到右依次挑选出res[0]、res[1]、res[2]，所以，如果str[5..9]中任何一个字符被挑选成为res[0]，之后过程是在挑选位置的右边继续挑选，那么一定会错过b字符，所以在str[0..4]上挑选res[0]。

3．在str[0..4]上找到字典序最小的字符，即str[1]=='a'，它就是res[0]。

4．在挑选字符 str[1]的右边，字符串为"acbaccac"，删掉所有的’a’字符变为"cbccc"，令str="cbccc"，下面找res[1]。

5．建立str[0..4]的字频统计，b有1个、c有4个。

6．从左往右遍历 str[0..4]，遍历到字符的字频统计减1，当发现某一种字符的字频统计已经为0时，遍历停止。当遍历完"cb"时，字频统计为b有0个、c有3个，发现b的字频已经为0，所以停止遍历，当前遍历到 str[1]。str[2..4]为"ccc"已经没有 b 了，所以如果 str[2..4]中任何一个字符被挑选成为res[1]，之后的过程是在挑选位置的右边继续挑选，那么一定会错过b字符，所以在str[0..1]上挑选res[1]。

7．在str[0..1]上找到字典序最小的字符，即str[1]=='b'，它就是res[1]。

8．在挑选字符str[1]的右边，字符串为"ccc"，删掉所有的’b’字符，仍为"ccc"，令str="ccc"，下面找res[2]。

9．建立str[0..2]的字频统计，c有3个。

10．从左往右遍历str[0..2]，遍历到字符的字频统计减1，当发现某一种字符的字频统计已经为0时，遍历停止。当遍历完"ccc"时，字频统计为c，有0个，当前遍历到str[2]。右边没有字符了，当然无法成为res[2]，所以在str[0..2]上挑选res[2]。

11．在str[0..2]上找到字典序最小的字符，即str[0]=='c'，它就是res[2]。整个过程结束。

如上过程虽然是用例子来说明的，但是整个过程其实比较简单。根据字频统计，遍历 str时找到一个前缀str[0..R]，然后在str[0..R]中找到最小ASCII码的字符str[X]，就是结果字符串的当前字符。然后令str=(str[X+1..R]去掉所有str[X]得到的字符串)，重复整个过程，找到结果字符串的下一个字符，直到res生成完毕。如果str长度为N，不同的字符有K种，每找到一个res[i]，都要重新建立字频统计以及在整个字符串中删除已经找到的字符，所以时间复杂度为O(K×N)。根据题目描述，str中全是小写字母，所以K不会超过26，则时间复杂度为O(N)。全部过程的代码实现请看如下removeDuplicateLetters方法。

public String removeDuplicateLetters(String s) {

char[] str = s.toCharArray();

// 小写字母ASCII码值范围为[97～122]，所以用长度为26的数组做次数统计

// 如果map[i] > -1，则代表ASCII码值为i的字符的出现次数

// 如果map[i] == -1，则代表ASCII码值为i的字符不

 



## 翻转字符串

【题目】

给定一个字符类型的数组chas，请在单词间做逆序调整。只要做到单词的顺序逆序即可，对空格的位置没有特别要求。

【举例】

如果把chas看作字符串为"dog loves pig"，调整成"pig Loves dog"。

如果把chas看作字符串为"I'm a student."，调整成"student. a I'm"。

补充问题：给定一个字符类型的数组chas和一个整数size，请把大小为size的左半区整体移到右半区，右半区整体移到左边。

【举例】

如果把chas看作字符串为"ABCDE", size=3，调整成"DEABC"。

【要求】

如果chas长度为N，两道题都要求时间复杂度为O(N)，额外空间复杂度为O(1)。

【解答】

原问题。首先把 chas 整体逆序。在逆序之后，遍历 chas 找到每一个单词，然后把每个单词里的字符逆序处理即可。比如“dog loves pig”，先整体逆序变为“gip sevol god”，然后每个单词进行逆序处理就变成了“pig loves dog”。逆序之后找每一个单词的逻辑，做到不出错即可。全部过程请参看如下代码中的rotateWord方法。

> public void rotateWord(char[] chas) {
>
> if (chas == null || chas.length == 0) {
>
> return;
>
> }
>
> reverse(chas, 0, chas.length - 1);
>
> int l = -1;
>
> int r = -1;
>
> for (int i = 0; i < chas.length; i++) {
>
> if (chas[i] ! = ' ') {
>
> l = i == 0 || chas[i - 1] == ' ' ? i : l;
>
> r = i == chas.length - 1 || chas[i + 1] == ' ' ? i : r;
>
> }
>
> if (l ! = -1 && r ! = -1) {
>
> reverse(chas, l, r);
>
> l = -1;
>
> r = -1;
>
> }
>
> }
>
> }
>
> public void reverse(char[] chas, int start, int end) {
>
> char tmp = 0;
>
> while (start < end) {
>
> tmp = chas[start];
>
> chas[start] = chas[end];
>
> chas[end] = tmp;
>
> start++;
>
> end--;
>
> }
>
> }

补充问题，方法一。先把chas[0..size-1]部分逆序处理，再把chas[size..N-1]部分逆序处理，最后把chas整体逆序处理即可。比如，chas="ABCDE", size=3。先把chas[0..2]部分逆序处理， chas变为"CBADE"，再把chas[3..4]部分逆序处理，chas变为"CBAED"，最后把chas整体逆序处理，chas变为"DEABC"。具体过程请参看如下代码中的rotate1方法。

public static void rotate1(char[] chas, int size) {



## 字符串的调整与替换

【题目】

给定一个字符类型的数组 chas[], chas 右半区全是空字符，左半区不含有空字符。现在想将左半区中所有的空格字符替换成"%20"，假设 chas 右半区足够大，可以满足替换所需要的空间，请完成替换函数。

【举例】

如果把chas的左半区看作字符串，为"a b c"，假设chas的右半区足够大。替换后，chas的左半区为"a%20b%20%20c"。

【要求】

替换函数的时间复杂度为O(N)，额外空间复杂度为O(1)。

补充问题：给定一个字符类型的数组chas[]，其中只含有数字字符和“*”字符。现在想把所有的“*”字符挪到chas的左边，数字字符挪到chas的右边。请完成调整函数。

【举例】

如果把chas看作字符串，为"12**345"。调整后chas为"**12345"。

【要求】

1．调整函数的时间复杂度为O(N)，额外空间复杂度为O(1)。

2．不得改变数字字符从左到右出现的顺序。

【难度】

士 ★☆☆☆

【解答】

原问题。遍历一遍可以得到两个信息，chas的左半区有多大，记为len，左半区的空格数有多少，记为 num，那么可知空格字符被“%20”替代后，长度将是 len+2×num。接下来从左半区的最后一个字符开始逆序遍历，同时将字符复制到新长度最后的位置，并依次向左逆序复制。遇到空格字符就依次对“0”、“2”和“%”进行复制。这样就可以得到替换后的chas数组。具体过程请参看如下代码中的replace方法。

> public void replace(char[] chas) {
>
> if (chas == null || chas.length == 0) {
>
> return;
>
> }
>
> int num = 0;
>
> int len = 0;
>
> for (len = 0; len < chas.length && chas[len] ! = 0; len++) {
>
> if (chas[len] == ' ') {
>
> num++;
>
> }
>
> }
>
> int j = len + num * 2 - 1;
>
> for (int i = len - 1; i > -1; i--) {
>
> if (chas[i] ! = ' ') {
>
> chas[j--] = chas[i];
>
> } else {
>
> chas[j--] = '0';
>
> chas[j--] = '2';
>
> chas[j--] = '%';
>
> }
>
> }
>
> }

补充问题。依然是从右向左逆序复制，遇到数字字符则直接复制，遇到“*”字符不复制。把数字字符复制完后，再把左半区全部设置成“*”即可。具体请参看如下代码中的modify方法。

> public void modify(char[] chas) {
>
> if (chas == null || chas.length == 0) {
>
> return;
>
> }
>
> int j = chas.length - 1;
>
> for (int i = chas.length - 1; i > -1; i--) {
>
> if (chas[i] ! = '*') {
>
> chas[j--] = chas[i];
>
>  2021-03-09 03:06:59



## 判断两个字符串是否互为变形词

【题目】

给定两个字符串str1和str2，如果str1和str2中出现的字符种类一样且每种字符出现的次数也一样，那么str1与str2互为变形词。请实现函数判断两个字符串是否互为变形词。

【举例】

str1="123", str2="231"，返回true。

str1="123", str2="2331"，返回false。

【解答】

如果字符串str1和str2长度不同，直接返回false。如果长度相同，假设出现字符的编码值在0～255之间，那么先申请一个长度为256的整型数组map, map[a]=b代表字符编码为a的字符出现了b次，初始时map[0..255]的值都是0。然后遍历字符串str1，统计每种字符出现的数量，比如遍历到字符’a'，其编码值为97，则令map[97]++。这样map就成了str1中每种字符的词频统计表。然后遍历字符串str2，每遍历到一个字符，都在map中把词频减下来，比如遍历到字符’a'，其编码值为97，则令map[97]--，如果减少之后的值小于0，直接返回false。如果遍历完str2, map中的值也没出现负值，则返回true。

具体请参看如下代码中的isDeformation方法。

> public boolean isDeformation(String str1, String str2) {
>
> if (str1 == null || str2 == null || str1.length() ! = str2.length()) {
>
> return false;
>
> }
>
> char[] chas1 = str1.toCharArray();
>
> char[] chas2 = str2.toCharArray();
>
> int[] map = new int[256];
>
> for (int i = 0; i < chas1.length; i++) {
>
> map[chas1[i]]++;
>
> }
>
> for (int i = 0; i < chas2.length; i++) {
>
> if (map[chas2[i]]-- == 0) {
>
> return false;
>
> }
>
> }
>
> return true;
>
> }
>
> 如果字符的类型有很多，可以用哈希表代替长度为256的整型数组，但整体过程不变。如果字符的种类为M, str1和str2的长度为N，那么该方法的时间复杂度为O(N)，额外空间复杂度为O(M)。

树

```

```



## start

年轻人总会找借口说这个东西不是我感兴趣的，所以做不好是应该的。但他们没有注意的是，你面对的事情中感兴趣的事情总是少数，这就使得大多数时候你做事情的态度总是很懈怠、很消极，这使你变成了一个懈怠的人。当你真正面对自己感兴趣的东西时，你发现你已经攥不紧拳头了

