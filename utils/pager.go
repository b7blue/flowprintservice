package utils

import (
	"fmt"
	"myflowprint/model"
	"regexp"
)

/* 	（1）当前数据一共有多少条。
（2）每页多少条，算出总页数。
（3）根据总页数情况，处理翻页链接。
（4）对页面上传入的 Get 或 Post 数据，需要从翻页链接中继续向后传。
（5）在页面显示时，根据每页数量和当前传入的页码，设置查询的 Limit 和 Skip，选择需要的数据。
（6）其他的操作，就是在 View 中显示翻页链接和数据列表的问题了。
*/

// 分页工具要被哪些页面使用：
// 1、会话展示 - detectsess_x
// 2、应用指纹管理 - train_info
// 3、单个应用指纹管理 - flowprints

/*
	处理顺序：
	1、根据是哪个表，传入的参数查数据库，找到总共有多少条数据
	2、根据maxlen与总数据量计算要分多少页pagenums
	3、每个单独的页上面，根据limit从相应的表中取数据，根据pagenums获得应该显示的页面链接

	传入参数：tablename, page
	返回值：data, urls,

*/

const maxlen int64 = 20
const urlnum int = 5

type Pager struct {
	Router    string
	URL       string
	Content   interface{}
	Firstpage bool
	Lastpage  bool
	PreURL    string
	NextURL   string
}

func (pager *Pager) get_round_page() {

	originurl := regexp.MustCompile(fmt.Sprintf(`^.*\?page=(\d+)$`, pager.Router))
	if !pager.Firstpage {
		if originurl.MatchString(pager.URL) {
			originurl.ReplaceAllString(pager.URL)
		}

	}
	if !islastpage {
		if baseurl.MatchString(c.Ctx.Request.RequestURI) {
			c.Data["nexturl"] = c.Ctx.Request.URL.Path + fmt.Sprintf("?page=%d", page+1)
		} else {
			c.Data["nexturl"] = c.Ctx.Request.RequestURI + fmt.Sprintf("&page=%d", page+1)
		}
	}
}

const limitlen int = 30

func Get_sess_page(detectid int, ip uint32, appname string, page int) (result []model.WebInfo, isfirstpage, islastpage bool) {
	if page <= 0 {
		return nil, true, true
	}
	datalen := model.Count_sess_by_term(detectid, ip, appname)
	if (page-1)*limitlen > datalen {
		return nil, true, true
	}
	result = model.Find_sess_by_term_page(detectid, ip, appname, page)
	if page*limitlen > datalen {
		islastpage = true
	}
	if page == 1 {
		isfirstpage = true
	}
	return
}
