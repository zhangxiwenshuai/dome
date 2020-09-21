package main

import (
	"fmt"
	"strings"
)

func main()  {
	x :="hello hello hello word"
	times := strings.Count(x,"hello")
	fmt.Printf("\n%s \n该字符串\"hello\"出现的字数为:%d",x,times)


}
