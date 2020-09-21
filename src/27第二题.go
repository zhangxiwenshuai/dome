package main

import "fmt"

func main() {
	var a int16 = 31
	var b int16 = 0
	a = 31
	b = 0
	b = a%8 + a/8%8*10
	fmt.Printf("十进制转八进制:")
	fmt.Printf("%d %d \n",a,b)

}


