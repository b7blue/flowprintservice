package controllers

import (
	"flowprintservice/utils"
	"fmt"
	"html/template"
	"log"
	"myflowprint/model"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type sess struct {
	ID      int
	Appname string
	Ip      string
	Port    uint16
	Packet  int
	Flow    int
	Start   string
	End     string
}

type SessDisplayController struct {
	beego.Controller
}

func (c *SessDisplayController) Get() {
	c.Data["IsDetecting"] = false
	c.Data["HaveDetected"] = false
	// 查看是否已经登陆
	tempid := c.Ctx.GetCookie("tempid")
	// 如果已经登陆，判断是否已经进行过app检测（redis存储tempid-detectid-status）
	if tempid != "" {
		if uid := utils.AlreadyLogin(tempid); uid != "" {
			// 假如用户已经进行过app检测，根据detectid查表，获得所有会话
			if detectid := utils.GetDetectId(uid); detectid != 0 {
				c.Data["HaveDetected"] = true

				// 如果正在检测中，显示那个转圈圈页面
				if utils.IsDetecting(detectid) {
					c.Data["isDetecting"] = true
				}
				page, _ := c.GetInt("page", 1)
				ip, appname := utils.IP2uint32(c.GetString("ip")), c.GetString("appname")
				result, isfirstpage, islastpage := utils.Get_sess_page(detectid, ip, appname, page)

				allSess = allSess[:num]
				c.Data["allSess"] = allSess
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
	c.TplName = "session-display.html"
}

func (c *SessDisplayController) Post() {
	// 进行app识别
	// 根据有没有上传文件判断检测方式
	// 新建一个识别
	// 改变redis中存储的detectid的状态
	// 用channel异步执行newdetect，传入detectid、是否有pcap文件参数
	// 外部进程根据是否上传有文件，选择两种调用monitor.CatchSess()的方式
	// 外部进程会话分析完毕后，将更新数据库
	// 改变redis中detectid的状态，新增uid与detectid对应关系
	// 查看是否已经登陆
	tempid := c.Ctx.GetCookie("tempid")
	// 如果已经登陆，开始新的检测
	if tempid != "" {
		if uid := utils.AlreadyLogin(tempid); uid != "" {
			// 新建一个detect
			detectid := model.NewDetect(time.Now())
			// 改变redis中存储的detectid的状态
			utils.DetectStart(detectid)
			// 是否离线分析
			offline := true
			// 判断是否上传文件
			f, _, err := c.GetFile("pcapfile")
			if err != nil {
				log.Fatal("getfile err ", err)
				offline = false
			} else {
				c.SaveToFile("pcapfile", fmt.Sprintf("detectdata/%d.pcap", detectid))
			}
			defer f.Close()
			// 将参数通过channel传给检测的goroutine，进行app检测
			dinfo := utils.Detectinfo{
				Uid:      uid,
				Detectid: detectid,
				Offline:  offline,
			}
			utils.NewDetectChan <- dinfo
			// 返回重新get该页面
			c.Get()

		} else {
			// 否则跳转登录界面
			c.Redirect("/login", 302)
		}
	} else {
		// 否则跳转登录界面
		c.Redirect("/login", 302)
	}

	// c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	// c.TplName = "session-display.html"

}
