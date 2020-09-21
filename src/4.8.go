package main

import "fmt"

type cpmpany_staff struct {
	name string
	cord int
	area string
	salary int
	tallage float32
}

func main()  {
	z := new(cpmpany_staff)
	z.name = "张大大"
	z.area = "未知"
	z.salary = 7890
	fmt.Printf("%#v",z)
	fmt.Printf("\n-------------\n")
	*z = zhang(*z)
	fmt.Printf("应收税%f",z.tallage)
}
func zhang(x cpmpany_staff)cpmpany_staff  {
	w:= x.salary
	if x.salary>5000 {
		x.tallage = float32(w-5000)*0.2
	}else {
		x.tallage = 0
	}
	return x
}