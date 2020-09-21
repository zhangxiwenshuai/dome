package main

import "fmt"

func main()  {
	a:=1939
	zhang(a)
	b:=2000
	zhang(b)
}
func zhang(c int)  {
	if c%4==0 && c%100==0 || c/400==0{
		fmt.Printf("%d该年是闰年\n",c)
	}else {
		fmt.Printf("%d该年不是闰年\n",c)
	}
}
