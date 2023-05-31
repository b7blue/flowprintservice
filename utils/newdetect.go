package utils

import (
	"fmt"
	"myflowprint/flowprintfactory"
	"myflowprint/monitor"
)

type Detectinfo struct {
	Uid      string
	Detectid int
	Offline  bool
}

var NewDetectChan = make(chan Detectinfo)

func init() {
	go func() {
		for {
			select {
			case dinfo := <-NewDetectChan:
				newdetect(dinfo)
			default:
			}
		}
	}()
}

func newdetect(dinfo Detectinfo) {
	uid := dinfo.Uid
	detectid := dinfo.Detectid
	offline := dinfo.Offline
	pcapfile := ""
	if offline {
		pcapfile = fmt.Sprintf("%d.pcap", detectid)
	}

	// 外部进程根据是否上传有文件，选择两种调用monitor.CatchSess()的方式
	monitor.CatchSess(false, detectid, pcapfile)
	flowprintfactory.Fingerprint(detectid, false)
	// 外部进程会话分析完毕后，将更新数据库，改变redis中detectid的状态
	DetectEnd(detectid)
	// 设置uid与detectid的对应关系
	SetDetectId(uid, detectid)

}
