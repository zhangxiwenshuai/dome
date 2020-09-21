package main

import(
	"fmt"
	"strings"
	)
func main(){
	var str string = "我爱你中国,I LOVE CHINA"
	strings.Split(str,"")
	fmt.Printf("长度%d\n",len(str))
	for _,value:=range str {
		fmt.Printf("%c  ",value)
	}

}