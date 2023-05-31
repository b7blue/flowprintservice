package utils

import (
	"context"
	"log"
	"os/exec"
)

var cancel context.CancelFunc
var cmd *exec.Cmd

func StartRedis() {
	var ctx context.Context
	ctx, cancel = context.WithCancel(context.Background())
	cmd = exec.CommandContext(ctx, "C:/Users/70408/Redis/redis-server.exe")
	if err := cmd.Start(); err != nil {
		log.Println("启动redis失败", err)
	} else {
		log.Println("启动redis成功")
	}

}

func StopRedis() {
	cancel()
	log.Println("关闭redis ", cmd.Wait())
}
