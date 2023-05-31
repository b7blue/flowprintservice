package controllers

import (
	"html/template"
	"myflowprint/model"

	beego "github.com/beego/beego/v2/server/web"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "register.html"
}

func (c *SignupController) Post() {
	// 从表单中获取新用户的信息
	uid, pw := c.GetString("uid"), c.GetString("pw")
	result := ""

	// 将新用户插入user表
	if model.NewUser(uid, pw) {
		result = "OK"
	} else {
		result = "该用户已存在"
	}

	c.Ctx.WriteString(result)

}
