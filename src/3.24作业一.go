package main

import(
"fmt"
"strings"
)

func main()  {
	str :="http://192.168.10.100:8080/Day33_Servlet/aa.jpeg"
	fmt.Printf("%s 文件路径:",str)
	z :=strings.LastIndex(str,"/")
	str2 :=  str[z+1:]
	z = strings.LastIndex(str2,".")
	str3 := str2[z+1:]
	str2 = str2[:z]
	fmt.Printf("\n该路径文件名为:%s 扩展名为:%s",str2,str3)

}
