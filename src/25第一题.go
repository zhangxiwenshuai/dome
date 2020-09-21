package main

import "fmt"

func main(){
	const(
		z1 int8 = iota
		z2
		z3
		z4
		z5
	)
  fmt.Println(z1,z2,z3,z4,z5)
}
