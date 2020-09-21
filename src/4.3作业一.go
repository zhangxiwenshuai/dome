package main

import "fmt"

func main()  {
	z := [5]int{11,22,33,44,55}
	P := &z
	fmt.Printf("数组的地址%p \nP的变量类型%T",&z,P)
	fmt.Printf("\n---------\n")
	var o [6]string = [6]string{"999","99","zxw","zz","xx","ww"}
	var p [6]*string
	for i:=0;i<len(o);i++  {
		p[i] = &o[i]
	}
	fmt.Printf("%p",p)
	*p[5]="dd"
	fmt.Printf("\n%s",o)
	fmt.Printf("\n---------\n")
	q := 11
	w := 22
	e := chang3(q,w)
	fmt.Printf("%d+%d=%d",q,w,*e)
}
func chang3(a int,q int) *int  {
	var w int = a+q
	return &w
}