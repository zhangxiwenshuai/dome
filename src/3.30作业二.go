package main

import "fmt"

func main()  {
	fmt.Printf("作业二\n")
	z:=10
	func (w int){
		x:=1
		for i:=1;i<=w;i++{
			x*=i
		}
		fmt.Printf("%d的阶乘为%d",w,x)
	}(z)
}
