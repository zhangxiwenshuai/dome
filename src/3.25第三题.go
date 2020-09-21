package main

import "fmt"

func main() {
	num:=8.1
	z:=x(num)
	fmt.Println("圆形的面积是：",z)
}
func x(n float64) float64{
	w:=3.14*(n*n)
	return w
}