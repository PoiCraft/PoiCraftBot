// Package syntax 实现SiS机器人支持的语法的解析和执行
package syntax

import (
	"regexp"
	"strings"

	"github.com/poicraft/bot/exec"
)

var (
	// 指令前缀，通常为cq码[CQ:at,qq=<机器人qq>]
	CmdPrefix string
)

var expMyID = regexp.MustCompile(`^\s*(?i)MyID\s*[=＝]\s*([0-9A-Za-z_]{3,16})\s*$`)

// GroupMsg 处理从游戏群接收到的消息，若为合法命令则进行相应的处理。并发安全
// 返回值指示是否拦截本消息
func GroupMsg(from int64, msg string, ret func(msg string)) bool {
	// 识别MyID指令
	// if match := expMyID.FindStringSubmatch(msg); len(match) == 2 {
	// 	whitelist.MyID(from, match[1], ret)
	// 	return true
	// }

	// 识别@指令
	if strings.HasPrefix(msg, CmdPrefix) {
		cmd := msg[len(CmdPrefix):]
		args := strings.Fields(cmd)
		if len(args) < 1 { // 如果没有首单词则不处理
			return false
		}

		switch args[0] {
		case "exec": // ping指令
			return exec.Exec(args, ret)
		}
	}

	return false
}
