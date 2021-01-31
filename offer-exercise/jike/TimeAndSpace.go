package jike

import "fmt"

//不同情况下的时间复杂度

func main() {

}

//O(1)
func Time1(n int) {
	i := 100
	fmt.Println("hey,your input is:", i)
}

//O(n)
func TimeN(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("hey,your input is", i)
	}
}

//O(n*2)
func TimeN2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println("your input is:", i, j)
		}
	}
}

//O(logn)
func TimeLogN(n int) {
	for i := 0; i < n; i = i * i {
		fmt.Println("your input is:", i)
	}
}

//O(n!)
func TimeN1(n int){
	for i:=0;i<factor(n);i++{
		fmt.Println("your input is:",i)
	}
}