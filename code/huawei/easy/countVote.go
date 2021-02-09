package main

import (
	"fmt"
	"strconv"
)

//输入候选人的人数，第二行输入候选人的名字，第三行输入投票人的人数，第四行输入投票。
//每行输出候选人的名字和得票数量。

//input:
//4
//A B C D
//8
//A B C D E F G H

//output:
//A : 1
//B : 1
//C : 1
//D : 1
//Invalid : 4

//func main() {
//	sc := bufio.NewScanner(os.Stdin)
//	for sc.Scan() {
//		candidateNum := sc.Text()
//		candidateNames := sc.Text()
//	}
//}

func main() {
	//candidateNum := 4
	candidateNames := []string{"A", "B", "C", "D"}
	//voteNum := 8
	voteNames := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	var candidateMap map[string]int
	candidateMap = make(map[string]int)
	for _, v := range candidateNames {
		candidateMap[v] = 0
	}
	var unsigned int
	for _, v := range voteNames {
		_, ok := candidateMap[v]
		if ok {
			candidateMap[v]++
		} else {
			unsigned++
		}
	}
	for k, v := range candidateMap {
		//fmt.Println(k, v)
		content := k + ":" + strconv.Itoa(v)
		fmt.Println(content)
	}
	if unsigned != 0 {
		fmt.Println("Invalid:", unsigned)
	}
	fmt.Println(candidateMap)
}
