package main

import "fmt"

func main()  {
	num1,num2,str := 0,0,0
	println("请输入两数之间运输符代数")
	print("\n 0|1++2-- 1|+ 2|- 3|* 4|/5% 例如输出4 2 1回车结果为4-1=3\n")
	fmt.Scanf("%d %d %d",&num1,&str,&num2)
	switch  str {
	case 0: num2--;num1++; fmt.Printf("\nnum1++%dnum2--%d",num1,num2)
	case 1: fmt.Printf("\n %d+%d=%d",num1,num2,num1+num2)
	case 2: fmt.Printf("\n %d-%d=%d",num1,num2,num1-num2)
	case 3: fmt.Printf("\n %d*%d=%d",num1,num2,num1*num2)
	case 4: fmt.Printf("\n %d/%d=%d",num1,num2,num1/num2)
	case 5: fmt.Printf("\n %d余%d=%d",num1,num2,num1%num2)
	default:
		println("您输入了错误的信息")
	}

}
