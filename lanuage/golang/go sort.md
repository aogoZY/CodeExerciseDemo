# go排序算法总结

### 冒泡

思想:比较相邻两个数,每次都将数字大的往后放.

步骤:

1. 比较相邻的元素。如果第一个比第二个大，就交换它们两个；
2. 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
3. 针对所有的元素重复以上的步骤，除了最后一个；
4. 重复步骤1~3，直到排序完成。

图:

![img](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015223238449-2146169197.gif)

```
func bubbleSort(input []int) []int {
   for i := 0; i < len(input); i++ {
      for j := 0; j < len(input)-i-1; j++ {
         if input[j] > input[j+1] {
            input[j], input[j+1] = input[j+1], input[j]
         }
      }
   }
   return input
}
```

### 冒泡优化

思想:外层循环定个flag,若是内层循环遍历一遍之后flag没变,说明此时数列已经有序了,可直接退出.

```
func bubbleSortV2(input []int) []int {
   for i := 0; i < len(input); i++ {
      flag := true
      for j := 0; j < len(input)-i-1; j++ {
         if input[j] > input[j+1] {
            input[j], input[j+1] = input[j+1], input[j]
            flag = false
         }

      }
      if flag == true {
         break
      }
   }
   return input
}
```

### 选择排序

思想:在未排序序列中找到最小元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。

步骤:

1. 初始状态：无序区为R[1..n]，有序区为空；

2. 第i趟排序(i=1,2,3…n-1)开始时，当前有序区和无序区分别为R[1..i-1]和R(i..n）。该趟排序从当前无序区中-选出关键字最小的记录 R[k]，将它与无序区的第1个记录R交换，使R[1..i]和R[i+1..n)分别变为记录个数增加1个的新有序区和记录个数减少1个的新无序区

3. n-1趟结束，数组有序化了。

   图:

   ![img](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015224719590-1433219824.gif)

   

```
//选择排序
//每次找到最小值 放到最小区间的末尾 循环
func selectSort(input []int) []int {
   for i := 0; i < len(input); i++ {
      index := i
      for j := i; j < len(input); j++ {
         if input[j] < input[index] {
            index = j
         }
      }
      input[i], input[index] = input[index], input[i]
   }
   return input
}
```

### 插入排序

思想:对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。

步骤:

1. 从第一个元素开始，该元素可以认为已经被排序；
2. 取出下一个元素，在已经排序的元素序列中从后向前扫描；
3. 如果该元素（已排序）大于新元素，将该元素移到下一位置；
4. 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
5. 将新元素插入到该位置后；
6. 重复步骤2~5。

图:

![img](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015225645277-1151100000.gif)

```
//插入排序
//将未排序的元素 和以排序的元素笔记 直到找到以排序元素小于未排序元素 将其插入后面
func insertSort(input []int) []int {
   for i := 1; i < len(input); i++ {
      preIndex := i - 1
      currentValue := input[i]
      for preIndex >= 0 && input[preIndex] > currentValue {
         input[preIndex+1] = input[preIndex]
         preIndex--
      }
      input[preIndex+1] = currentValue
   }
   return input
}
```

### 快排

思想:通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

1. 从数列中挑出一个元素，称为 “基准”（pivot）；
2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。

图:

![使用Golang实现的快速排序](http://static.open-open.com/lib/uploadImg/20140908/20140908184909_185.jpg)

```
//快排 分成多个子串来做处理 对每个子串 随机选取一个标准值 将小于标准值的数置于左边 将大的标准值置于右边
func quitSort(input []int, left, right int) []int {
   if left < right {
      pos := partion(input, left, right)
      fmt.Println(pos)
      quitSort(input, left, pos-1)
      quitSort(input, pos+1, right)
   }
   return input
}

//返回标志值索引位置
func partion(arr []int, left, right int) int {
   value := arr[right]
   i := left - 1
   for j := left; j < right; j++ {
      if arr[j] <= value {
         arr[i+1], arr[j] = arr[j], arr[i+1]
         i++
      }
   }
   arr[i+1], arr[right] = arr[right], arr[i+1]
   fmt.Println(arr)
   return i + 1
}
```

### 归并

思想:分治法（Divide and Conquer）典型应用.将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。

步骤:

1. 把长度为n的输入序列分成两个长度为n/2的子序列；
2. 对这两个子序列分别采用归并排序；
3. 将两个排序好的子序列合并成一个最终的排序序列。

图:

![](https://images2015.cnblogs.com/blog/1024555/201612/1024555-20161218163120151-452283750.png)

![img](https://images2015.cnblogs.com/blog/1024555/201612/1024555-20161218194508761-468169540.png)



```
//归并 先对半分 将其子串排序
func mergeSort(arr []int) []int {
   if len(arr) < 2 {
      return arr
   }
   mid := len(arr) / 2
   left := mergeSort(arr[:mid])
   right := mergeSort(arr[mid:])
   return merge(left, right)
}

func merge(left []int, right [] int) []int {
   leftLength := len(left)
   rightLength := len(right)
   var res []int
   var li, ri int
   for li < leftLength && ri < rightLength {
      if left[li] < right[ri] {
         res = append(res, left[li])
         li++
      } else {
         res = append(res, right[ri])
         ri++
      }
   }

   if li < leftLength {
      res = append(res, left[li:]...)
   } else {
      res = append(res, right[ri:]...)
   }
   return res
}
```

# 稳定性

- 不稳定：快速排序、希尔排序、堆排序、直接选择排序

- 稳定：基数排序、冒泡排序、直接插入排序、折半插入排序、归并排序


