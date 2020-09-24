package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "1571021611@qq.com"
	c.TplName = "index.tpl"
}

func (c * MainController) post(){
	for i:=0;i<10;i++  {
		fmt.Printf("第%d次打印\n",i)
	}
}
func (c *MainController) post(){
	name := c.Ctx.Request.FormValue("name")
	age := c.Ctx.Request.FormValue("age")
	sex := c.Ctx.Request.FormValue("sex")
	fmt.Println(name,age,sex)

	if name != "admin" && age != "18"{
		c.Ctx.WriteString("数据效验失败")
		return
	}
}