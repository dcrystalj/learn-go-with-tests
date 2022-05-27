package iteration

import "strings"

func Repeat(s string, r int) string {
	repeated := ""
	for i := 0; i < r; i += 1 {
		repeated += s
	}
	return repeated
}

func Repeat2(s string, r int) string {
	return strings.Repeat(s, r)
}
