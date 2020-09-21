package main

import (
	"errors"
	"fmt"
)

type Person struct {
	name string
	age int
	address string
}

func JudgeAge(per Person) (bool,error)  {
	if per.age>=18{
		return true,nil
	}else if 0<per.age {
		return false,nil
	}else {
		var err error =errors.New("年龄不合法")
		return false,err
	}
}
func main()  {
	z := Person{"李小花",0,"江西"}
	fmt.Println(z)
	x,w := JudgeAge(z)
	if w!=nil {
		fmt.Printf("%s",w)
	}else {
		if x==true {
			fmt.Printf("已成年")
		}else {
			fmt.Printf("未成年")
		}
	}
	fmt.Printf("\n--------------------\n")
	d := [6]int{1,2,3,4,5,6}
	for i:=0;i<10;i++  {
		if i>=len(d) {
			panic("超出数组长度")
		}
	}
}
