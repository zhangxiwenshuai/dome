package main

import "fmt"

func main() {
	fmt.Printf("作业一\n")
	n := 9
	sun(n, n-1, 1)
}
func sun(n int,w int,s int)  {
	for i:=1;i<=s;i++{
		fmt.Printf("%d",i)
	}
	fmt.Printf(" ")
	if w>0 {
		s++
		sun(n,w-1,s)
	}

}
