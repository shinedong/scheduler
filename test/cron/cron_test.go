package cron

import (
	"github.com/gorhill/cronexpr"
	"testing"
	"time"
)

func TestCron(t *testing.T) {

	var cronExpr string = "*/5 * * * * * *" // 每隔5秒执行一次, 秒(0-59) 分(0-59) 时(0-23) 天(1-31) 月(1-12) 星期(0-6) 年(?-2099)
	expression, err := cronexpr.Parse(cronExpr)
	if err != nil {
		t.Fatalf("解析失败")
	}
	now := time.Now() // 当前时间
	t.Log("当前时间=", now.String())
	nextTime := expression.Next(now)
	// 等待定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		t.Log("调度时间=", nextTime.String())
	})

	time.Sleep(5 * time.Second)
}
