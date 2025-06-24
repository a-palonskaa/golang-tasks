//go:build !solution

package speller

import (
	"strings"
)

type scale struct {
	value int64
	name  string
}

var scales = []scale{
	{1e9, "billion"},
	{1e6, "million"},
	{1e3, "thousand"},
	{1, ""},
}

var (
	ones = map[int]string{1: "one", 2: "two", 3: "three", 4: "four",
		5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}
	teens = map[int]string{10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen"}
	tens = map[int]string{2: "twenty", 3: "thirty", 4: "forty",
		5: "fifty", 6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety"}
)

var MaxWords = 12

func Spell(num int64) string {
	if num == 0 {
		return "zero"
	}

	parts := make([]string, 0, MaxWords)

	if num < 0 {
		parts = append(parts, "minus")
		num = -num
	}

	convertChunk := func(n int) {
		if n >= 100 {
			parts = append(parts, ones[n/100], "hundred")
			n %= 100
		}

		switch {
		case n == 0:
		case n < 10:
			parts = append(parts, ones[n])
		case n < 20:
			parts = append(parts, teens[n])
		default:
			word := tens[n/10]
			if n%10 != 0 {
				word += "-" + ones[n%10]
			}
			parts = append(parts, word)
		}
	}

	for _, s := range scales {
		if num >= s.value {
			chunk := int(num / s.value)
			num %= s.value

			convertChunk(chunk)

			if s.name != "" {
				parts = append(parts, s.name)
			}
		}
	}
	return strings.Join(parts, " ")
}
