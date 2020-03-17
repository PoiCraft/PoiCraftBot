package main

import (
	"fmt"

	"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"
	"github.com/poicraft/bot/data"
	"github.com/poicraft/bot/log"
	"github.com/poicraft/bot/syntax"
)

//go:generate cqcfg -c .
// cqp: 名称: PoiCtartBot
// cqp: 版本: 1.0.2:0
// cqp: 作者: topjohncian
// cqp: 简介: PoiCraftBot
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "com.poicraft.bot" // TODO: 修改为这个插件的ID
	cqp.Enable = onStart
	cqp.Disable = onStop
	cqp.Exit = onStop
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg
	data.Logger = log.NewLogger("Data")
}

var Logger = log.NewLogger("Main")

// 插件生命周期开始
func onStart() int32 {
	// 连接数据源
	err := data.Init(cqp.GetAppDir())
	if err != nil {
		Logger.Errorf("初始化数据源失败: %v", err)
	}

	// 将登录账号载入命令解析器（用于识别@）
	syntax.CmdPrefix = fmt.Sprintf("[CQ:at,qq=%d]", cqp.GetLoginQQ())

	return 0
}

// 插件生命周期结束
func onStop() int32 {
	err := data.Close()
	if err != nil {
		Logger.Errorf("释放数据源失败: %v", err)
	}
	return 0
}
func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return Ignore
}

// 群消息事件
func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	if fromQQ == 80000000 { // 忽略匿名
		return Ignore
	}

	ret := func(resp string) {
		cqp.SendGroupMsg(fromGroup, resp)
	}
	if syntax.GroupMsg(fromQQ, msg, ret) {
		return Intercept
	}
	return Ignore
}

const (
	Ignore    int32 = 0 //忽略消息
	Intercept       = 1 //拦截消息

	Allow = 1 // 允许进群
	Deny  = 2 // 拒绝进群
)
