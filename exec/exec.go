package exec

import "strings"

func Exec(args []string, ret func(msg string)) bool {
	arg := strings.Join(args, ` `)
	arg = strings.Replace(arg, args[0]+` `, ``, 1)

	ret(arg)
	return true
}
