package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	write := http.DefaultClient
	g, err := write.Get("http://127.0.0.110:9001/userlogin")
	if err != nil {
		fmt.Printf("请求失败 %s\n", err)
		return
	}
	z, w := ioutil.ReadAll(g.Body)
	if w != nil {
		fmt.Printf("读取失败 %s\n", w)
		return
	}
	g.Body.Close()
	fmt.Printf("结果:%s", z)
}