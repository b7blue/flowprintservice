package controllers

import (
	"flowprintservice/utils"
	"html/template"
	"log"
	"myflowprint/model"

	beego "github.com/beego/beego/v2/server/web"
)

type flowprint struct {
	Appname string
	Pid     int
	Dip     string
	Dport   uint16
}

type ManageOnePrintController struct {
	beego.Controller
}

func (c *ManageOnePrintController) Get() {
	// 查看是否已经登陆
	tempid := c.Ctx.GetCookie("tempid")
	// 如果已经登陆
	if tempid != "" {
		if uid := utils.AlreadyLogin(tempid); uid != "" {
			id, err := c.GetInt("id")
			// 编辑单个指纹的链接是否合法？是否有pid参数
			if err == nil && id != 0 {
				// pid是否合法，对应的app是否已经生成指纹？
				if appname := model.IsAppPrintsExist(id); appname != "" {
					rawappfp := model.GetAppPrints(appname)
					appfp := make([]flowprint, len(rawappfp))
					for i, fp := range rawappfp {
						appfp[i] = flowprint{
							Appname: fp.Appname,
							Pid:     fp.Pid,
							Dip:     utils.FormatIP(fp.Dip),
							Dport:   fp.Dport,
						}
					}
					c.Data["appfp"] = appfp
					c.Data["appname"] = appname
					c.Data["fpnum"] = model.GetPrintsNum(appname)

				} else {
					// pid参数不合法，跳转指纹管理主页
					c.Redirect("/manageprints", 302)
				}

			} else {
				// 链接不合法，跳转指纹管理主页
				c.Redirect("/manageprints", 302)
			}

		} else {
			// 否则跳转登录界面
			c.Redirect("/login", 302)
		}

	} else {
		// 否则跳转登录界面
		c.Redirect("/login", 302)
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "manage-oneprint.html"
}

func (c *ManageOnePrintController) Post() {
	// 查看是否已经登陆
	tempid := c.Ctx.GetCookie("tempid")
	// 如果已经登陆
	if tempid != "" {
		if uid := utils.AlreadyLogin(tempid); uid != "" {
			id, err := c.GetInt("id")
			// 编辑单个指纹的链接是否合法？是否有pid参数
			if err == nil && id != 0 {
				// pid是否合法，对应的app是否已经生成指纹？
				if appname := model.IsAppPrintsExist(id); appname != "" {
					ip, oldip := utils.IP2uint32(c.GetString("ip")), utils.IP2uint32(c.GetString("oldip"))
					port, _ := c.GetInt("port")
					oldport, _ := c.GetInt("oldport")
					op := c.GetString("op")
					appname = c.GetString("appname")
					pid, _ := c.GetInt("pid")

					if op == "del" {
						if err := model.DelFingerprint(appname, pid, oldip, uint16(oldport)); err != nil {
							log.Println("删除网络目的地失败", err, "原内容为：", oldip, oldport)
						} else {
							log.Println("删除网络目的地成功，原内容为：", oldip, oldport)
						}

					} else if op == "edit" {
						if err := model.UpdateFingerprint(appname, pid, oldip, ip, uint16(oldport), uint16(port)); err != nil {
							log.Println("更新网络目的地失败", err, "原内容为：", oldip, oldport)
						} else {
							log.Println("更新网络目的地成功，新内容为:", ip, port)
						}
					} else {
						// op不合法
						c.Redirect("/manageprints", 302)
					}
					// 从数据库读取最新的指纹列表
					appfp := model.GetAppPrints(appname)
					c.Data["appfp"] = appfp
					c.Data["appname"] = appname
					c.Data["fpnum"] = model.GetPrintsNum(appname)

				} else {
					// pid参数不合法，跳转指纹管理主页
					c.Redirect("/manageprints", 302)
				}

			} else {
				// 链接不合法，跳转指纹管理主页
				c.Redirect("/manageprints", 302)
			}

		} else {
			// 否则跳转登录界面
			c.Redirect("/login", 302)
		}

	} else {
		// 否则跳转登录界面
		c.Redirect("/login", 302)
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "manage-oneprint.html"
}
