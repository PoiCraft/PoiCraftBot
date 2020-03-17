package exec

import (
	"github.com/poicraft/bot/data"
	"strings"
)

func Exec(args []string, ret func(msg string)) bool {
	arg := strings.Join(args, ` `)
	arg = strings.Replace(arg, args[0]+` `, ``, 1)
	data.Logger.Info("执行：" + arg)
	ret(data.WebSocketExec(arg))
	return true
}
