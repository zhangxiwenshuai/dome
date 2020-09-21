package main

import "fmt"
func main()  {
	const(
	 zhang1 string = "浮点型变量"
     zhang2 string = "整型变量"
     zhang3 string = "字符型变量"
     zhang11 float32 = 1.0
	 zhang22 int = 1
	 zhang33 string = "张希文"
	)

    fmt.Printf("%s %T %s %T %s %T",zhang1,zhang11,zhang2,zhang22,zhang3,zhang33)

}
