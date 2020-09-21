package main

import "fmt"

func main()  {
	fmt.Printf("\n作业一:\n")
	a := []int{1,2,3,4,5}
	for _,value := range a{
		fmt.Printf("%d",value)
	}
	fmt.Printf("\n作业二:\n")
	str := []string{"小狗","小猫","小猪"}
	for _,value := range  str{
		fmt.Printf("%s",value)
	}
	fmt.Printf("\n作业三:\n")
	fol := []float32{1.11,2.22,3.33,4.44}
	var  F float32=0
	for _,value := range fol {
		F +=value
	}
	fmt.Printf("%0.2 总和%.2f",fol,F)
	fmt.Printf("\n作业四:\n")
	u8 :=[10]int{}
	for i:=0;i<10;i++{
		fmt.Scanf("%d",&u8[i])
	}
	fmt.Printf("%d",u8)
}