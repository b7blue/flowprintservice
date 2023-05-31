package controllers

import (
	"flowprintservice/utils"
	"html/template"
	"time"

	"myflowprint/model"

	beego "github.com/beego/beego/v2/server/web"
)

// type user struct {
// 	Id       int    `form:"-"`
// 	Email    string `form:"email"`
// 	Password string `form:"pwd"`
// }

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	result := ""
	uid, pw := c.GetString("uid"), c.GetString("pw")

	// 检验密码正确性
	if model.CheckPw(uid, pw) {
		// 假如密码正确则成功登录，生成tempid，设置cookie
		tempid := utils.SetTempId(uid)
		c.Ctx.SetCookie("tempid", tempid, time.Hour*24*7)
		result = "OK"
	} else {
		result = "用户名或密码错误"
	}
	c.Ctx.WriteString(result)

}
