package controllers

import (
	"flowprintservice/utils"
	"html/template"
	"myflowprint/model"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type ManagePrintsController struct {
	beego.Controller
}

func (c *ManagePrintsController) Get() {
	// 查看是否已经登陆
	tempid := c.Ctx.GetCookie("tempid")
	// 如果已经登陆
	if tempid != "" {
		if uid := utils.AlreadyLogin(tempid); uid != "" {
			appname := c.GetString("appname")
			applist := model.GetFLowprintAppList()
			if appname != "" {
				sublist := make([]model.TrainInfo, len(applist))
				num := 0
				for _, a := range applist {
					if strings.ContainsAny(a.Appname, appname) {
						sublist[num] = a
						num++
					}
				}
				applist = sublist[:num]
			}

			c.Data["applist"] = applist
		} else {
			// 否则跳转登录界面
			c.Redirect("/login", 302)
		}

	} else {
		// 否则跳转登录界面
		c.Redirect("/login", 302)
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "manage-prints.html"
}

func (c *ManagePrintsController) Post() {
	// 查看是否已经登陆
	tempid := c.Ctx.GetCookie("tempid")
	// 如果已经登陆
	if tempid != "" {
		if uid := utils.AlreadyLogin(tempid); uid != "" {

			appname := c.GetString("appname")
			appid, _ := c.GetInt("appid")
			// 检查appname和id的对应关系是否正确
			if rightname := model.IsAppPrintsExist(appid); rightname != "" && rightname == appname {
				// 删除app的指纹：在flowprints表中删，在trainlist中删，printinfo中删
				model.DelAppPrints(appname)
				model.PrintsDel(appid)
				model.DelPrintsInfo(appname)
			}

			// 获得最新applist
			applist := model.GetFLowprintAppList()
			c.Data["applist"] = applist
		} else {
			// 否则跳转登录界面
			c.Redirect("/login", 302)
		}

	} else {
		// 否则跳转登录界面
		c.Redirect("/login", 302)
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "manage-prints.html"
}
