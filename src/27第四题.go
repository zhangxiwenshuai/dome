package main

import "fmt"

func main()  {
    var a,b int
    a = 4
    b = 3
    res1 :=a<b&&b/2==1&&a%3!=0
    res2 :=(a+b)*3<a<<2||(a-b)>0

    fmt.Printf("表达式的值是:\n")
    fmt.Println("res1:=a<b&&b/2==1&&a余3!=0 %t \n",res1)
    fmt.Println("res2:=(a+b)*3<a<<2||(a-b)>0 %t \n",res2)
}