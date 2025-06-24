//go:build !solution

package varjoin

import "strings"

func Join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}

	totalLen := 0
	for _, s := range args {
		totalLen += len(s)
	}

	str := strings.Builder{}
	if sep == "" {
		str.Grow(totalLen)
		for _, val := range args {
			str.WriteString(val)
		}
	} else {
		totalLen += len(sep) * (len(args) - 1)
		str.Grow(totalLen)

		str.WriteString(args[0])
		for _, val := range args[1:] {
			str.WriteString(sep)
			str.WriteString(val)
		}
	}
	return str.String()
}
