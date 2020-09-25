package controllers

import (
	"BeegoHello01/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

/**
 *该结构体用于处理/register请求
 */
type RegisterContreller struct {
	beego.Controller
}

/**
*该方法用于处理post类型请求
 */

func (r *RegisterContreller) Post(){
	fmt.Println(r == nil)
	fmt.Println(r.Ctx == nil)
	fmt.Println(r.Ctx.Request == nil)
    //1、接收前端传递的数据
	body,err :=r.Ctx.GetBody()
	if err != nil{
		r.Ctx.WriteString("数据接收错误")
		return
	}
    bodyBytes,err :=ioutil.ReadAll(body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	var user models.User
    err = json.Unmarshal(bodyBytes,&user)
	if err != nil {
		r.Ctx.WriteString("数据解析错误请重试")
		return
	}
	fmt.Println(user.User)
    fmt.Println(user.Nick)
    r.Ctx.WriteString("数据接收成功，并成功解析")
}