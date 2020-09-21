package main

import "fmt"

func main()  {
	month :=0
	println("请输入1-12月份:")
	fmt.Scanf("%d",&month)
	switch  month {
	case 1: fmt.Printf("\n冬季")
	case 2: fmt.Printf("\n冬季")
	case 3: fmt.Printf("\n春季")
	case 4: fmt.Printf("\n春季")
	case 5: fmt.Printf("\n春季")
	case 6: fmt.Printf("\n夏季")
	case 7: fmt.Printf("\n夏季")
	case 8: fmt.Printf("\n夏季")
	case 9: fmt.Printf("\n秋季")
	case 10: fmt.Printf("\n秋季")
	case 11: fmt.Printf("\n秋季")
	case 12: fmt.Printf("\n冬季")
	default:
		println("您输入了错误的信息")
	}
    print("\n")

}