package main

import "fmt"

func main()  {
	a := 0
	println("请输入1-7的数值")
	fmt.Scanf("%d",&a)
	switch a {
	case 1: fmt.Printf("\n星期一")
	case 2: fmt.Printf("\n星期二")
	case 3: fmt.Printf("\n星期三")
	case 4: fmt.Printf("\n星期四")
	case 5: fmt.Printf("\n星期五")
	case 6: fmt.Printf("\n星期六")
	case 7: fmt.Printf("\n星期日")
	default:
		println("您输入了错误的信息")
	}
	print("\n")


}