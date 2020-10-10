package main
import "fmt"
func main() {
	for {
		var str string
		_,err := fmt.Scan(&str)
		if err != nil {
			return
		}
		// input your code
		str1 := ""
		for i := 0; i <len(str); i++ {
			if str[i] < '0' || str[i] > '9' {
				str1+=string(str[i])
			}else {
				if i == 0 || (str[i-1] < '0' || str[i-1] > '9'){
					str1 += "*"
				}
				str1 += string(str[i])
				if i == len(str) -1 || (str[i+1] < '0' || str[i+1] > '9') {
					str1 += "*"
				}
			}
		}
		fmt.Println(str1)
	}
}