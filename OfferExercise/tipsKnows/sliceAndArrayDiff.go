package main

import "fmt"

//slice和array的区别

func main()  {
	//定义slice切片 指针传递 改变原值
	slice:= []int{1,2,3}
	sliceCopy := slice
	sliceCopy[0]=9
	fmt.Println(slice)
	//9 2 3]    切片是指针 将原有值也改变了
	fmt.Println(sliceCopy)
	//[9 2 3]

	//定义array数组 值传递
	array:=[3]int{1,2,3}
	arrayCopy:=array    //数组是将整个复制的值传递
	arrayCopy[0]=9
	fmt.Println(array)
	//[1 2 3]
	fmt.Println(arrayCopy)
	//[9 2 3]

}
