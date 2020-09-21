package main

import (
	"net/http"
)
type num struct {
}

func (message *num)ServeHTTP(me http.ResponseWriter,mess *http.Request){
	m :=mess.URL
	if m.Path=="/userinfo"{
		me.Write([]byte("查询用户信息"))
		return
	}else if  m.Path=="/student"{
		me.Write([]byte("查询学生信息"))
	}else{
		http.NotFound(me,mess)
	}
	return
}

func main(){
	nua := num{}
	http.ListenAndServe("127.0.0.110:8090",&nua)
}
