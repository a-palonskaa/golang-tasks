//go:build !solution

package speller

import (
	"strings"
)

var (
	ones = []string{"", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine"}
	teens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens = []string{"", "", "twenty", "thirty", "forty",
		"fifty", "sixty", "seventy", "eighty", "ninety"}
)

func Spell(num int64) string {
	if num == 0 {
		return "zero"
	}

	var str strings.Builder

	if num < 0 {
		str.WriteString("minus")
		num = -num
	}

	convertLessThanHundred := func(str *strings.Builder, n int) {
		if n == 0 {
			return
		}

		switch {
		case n < 10:
			str.WriteString(ones[n])
		case n < 20:
			str.WriteString(teens[n%10])
		default:
			str.WriteString(tens[n/10])
			if n%10 != 0 {
				str.WriteString("-")
				str.WriteString(ones[n%10])
			}
		}
	}

	convert := func(str *strings.Builder, n int) {
		if n >= 100 {
			str.WriteString(ones[n/100])
			str.WriteString(" hundred")
			if remainder := n % 100; remainder > 0 {
				str.WriteString(" ")
				convertLessThanHundred(str, remainder)
			}
		} else {
			convertLessThanHundred(str, n)
		}
	}

	type scale struct {
		value int64
		name  string
	}

	scales := []scale{
		{1e9, "billion"},
		{1e6, "million"},
		{1e3, "thousand"},
		{1, ""},
	}

	for _, s := range scales {
		if num >= s.value {
			chunk := int(num / s.value)
			num %= s.value

			if str.Len() > 0 {
				str.WriteString(" ")
			}

			convert(&str, chunk)

			if s.name != "" {
				str.WriteString(" ")
				str.WriteString(s.name)
			}
		}
	}

	return str.String()
}
