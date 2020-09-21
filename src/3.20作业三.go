package main

import (
	"fmt"
	"strings"
)

func main()  {
	str :="WUHANJIAYOUZHONGGUOJIAYOU"
	var x bool = strings.Contains(str,"ZHONGGUO")
	if x {
		fmt.Printf("存在 位置为%d",strings.Index(str,"ZHONGGUO"))
	}else {
		fmt.Printf("不存在 ZHONGGUO")
	}

}
