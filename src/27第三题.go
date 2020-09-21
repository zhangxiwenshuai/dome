package main

import "fmt"

func main()  {
	var a,b int
	a = 5
	b = 6
	fmt.Printf("运算符运用:")
	fmt.Printf("a>b=%t a<b=%t a>=b=%t a<=b%t a==b=%t a!=b=%t \n",a>b,b>a,a>=b,a<=b,a==b,a!=b)

}