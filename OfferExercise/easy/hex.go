package main

import ("fmt"
	"io"
)

func main() {
	var l,r,m int
	for {
		_, err := fmt.Scanf("%d%d%d",&l,&r, &m)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			var hexList []int
			for i:=l;i<r+1;i++{
				binary:=hexToBinary(i,m)
				if binary != 0{
					hexList=append(hexList, binary)
				}
			}
			fmt.Println(len(hexList))
		}
	}
}


//func main(){
//	var l,r,m int
//	fmt.Scanln(&l)
//	fmt.Scanln(&r)
//	fmt.Scanln(&m)
//	var hexList []int
//	for i:=l+1;i<r;i++{
//		binary:=hexToBinary(i,m)
//		if binary != 0{
//			hexList=append(hexList, binary)
//		}
//	}
//	// fmt.Println(hexList)
//	fmt.Println(len(hexList))
//
//}

func hexToBinary(i,m int)(int){
	iCopy := i
	var count int
	for i>0{
		if i%2 == 1{
			count++
		}
		i=i/2
	}
	if count == m{
		return iCopy
	}else{
		return 0
	}

}
