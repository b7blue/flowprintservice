package main

import (
	_ "flowprintservice/routers"
	"flowprintservice/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 开启redis服务
	utils.StartRedis()
	defer utils.StopRedis()
	// 设置静态文件处理
	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
