package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/userlogin",func(pig http.ResponseWriter,dog *http.Request){
		fmt.Println("已发送回复")
		pig.Write([]byte("欢迎访问用户登录功能"))
	})
	zxw := http.ListenAndServe("127.1.1.111:9001",nil)
	if zxw!= nil{
		fmt.Println(zxw)
	}

}