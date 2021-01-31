package main

import "fmt"

func exist(board [][]string, word string) bool {
	for _,str :=range word{
		for _,i :=board[0]{
			for _,j:=board[]
		}
	}

}

func main()  {
	board :=[][]string{
		{"A","B","C","E"},
		{"S","F","C","S"},
		{"A","D","E","E"},
	}
	res := exist(board,"ABCCED")
	fmt.Println(res)
}
