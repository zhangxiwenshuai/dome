package main

import "fmt"

func main()  {
	a := []int{1,2,3,4,5}
	fmt.Printf("%d",a)
	a = zhang(a)
	fmt.Printf("%d",a)

}
func zhang(x []int) []int{
	w := len(x)
	for i:=0;i<w/2;i++  {
		x[i],x[w-i-1] = x[w-i-1],x[i]
	}
	return x
}