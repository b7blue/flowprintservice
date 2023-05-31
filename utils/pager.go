package utils

import (
	"myflowprint/model"
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
	tablename string
	page      int
	datanums  int64
	pagenums  int64

	Data interface{}
	Pre  []int
	Next []int
}

// 根据传回的pager是否为nil判断page合法与否
func NewPager(tablename string, page int) *Pager {
	datanums := model.Count(tablename)
	pagenums := (datanums + maxlen - 1) / maxlen

	// 判断传入的page是否合法
	if page == 0 {
		// 说明要查的是第一页
		page = 1
	} else if page < 0 || int64(page) > pagenums {
		return nil
	}

	return &Pager{
		tablename: tablename,
		page:      page,
		datanums:  datanums,
		pagenums:  pagenums,
	}
}

func (pager *Pager) Page() {
	pager.getpre()
	pager.getnext()
}

func (pager *Pager) getpre() {
	var l int
	if pager.page > urlnum {
		l = urlnum

	} else {
		l = pager.page - 1
	}
	pre := make([]int, l)
	for i := 1; i <= l; i++ {
		pre[l-i] = pager.page - i
	}
	pager.Pre = pre
}

func (pager *Pager) getnext() {
	var l int
	if int(pager.pagenums)-pager.page > urlnum {
		l = urlnum

	} else {
		l = int(pager.pagenums) - pager.page
	}
	next := make([]int, l)
	for i := 1; i <= l; i++ {
		next[i-1] = pager.page + 1
	}
	pager.Next = next
}

func (pager *Pager) getdata() {
	// 计算要查询的数据库范围
	
}
