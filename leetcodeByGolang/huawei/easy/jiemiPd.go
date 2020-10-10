package main

import (
"fmt"
)

func jiami(c byte) byte{

	if c>='A' && c<= 'Y'{
		c=c+'a'-'A'+1
		return c
	}else if c=='Z'{
		c='a'
		return c
	}

	switch c{
	case 'a','b','c':
		c='2'
	case 'd','e','f':
		c='3'
	case 'g','h','i':
		c='4'
	case 'j','k','l':
		c='5'
	case 'm','n','o':
		c='6'
	case 'p','q','r','s':
		c='7'
	case 't','u','v':
		c='8'
	case 'w','x','y','z':
		c='9'
	}

	return c
}

func main()  {
	var input []byte
	fmt.Scanf("%s",&input)
	for i,v := range input{
		input[i]=jiami(v)
	}
	fmt.Printf("%s",input)

}
