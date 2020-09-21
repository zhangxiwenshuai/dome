package main

import "shcoll"

func main()  {
	results :=Shcoll.Athletes{
		Name:"二狗子",
		Gender: "男",
		SportEvent:"跳高",
		Achi: 2,
	}
	results.Run()
}