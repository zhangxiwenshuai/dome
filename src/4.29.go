package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		zhang := request.FormValue("zhang")
		fmt.Println(zhang)
		password := request.Form.Get("password")
		fmt.Println(password)
		if zhang == "hello" && password == "123456" {
			writer.Write([]byte("恭喜登录成功"))
		}else{
			writer.Write([]byte("对不起，您的账号或者密码不正确，请重新尝试"))
		}
	})
	fmt.Println("启动监听服务")
	err:=http.ListenAndServe("127.0.0.1:8080",nil)
	if err!=nil {
		fmt.Println(err.Error())
	}
}
