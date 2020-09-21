package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "zhangxiwenshidashuaige"
	dage(str)
}
func dage(str string) {
	for a:='a';a<='z' ;a++  {
		x:=string(a)
		z:=strings.Count(str,x)
		if z!= 0{
			fmt.Println("该字符中",x,"出现了",z,"次")
		}
	}
}