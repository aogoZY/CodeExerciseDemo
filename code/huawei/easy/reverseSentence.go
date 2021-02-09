package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ResevieSentence(str string) {
	res :=strings.Split(str," ")
	for i:=len(res)-1;i>0;i--{
		fmt.Print((res[i])+" ")
	}
	fmt.Println(res[0])
	fmt.Println()
}

func main() {
	inputReader:=bufio.NewReader(os.Stdin)
	str,err:=inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	strRemove:=str[:len(str)-1]
	ResevieSentence(strRemove)
}