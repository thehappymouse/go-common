package utils

import (
	"bytes"
	"regexp"
)

// 根据正则，返回匹配的字符串
func ExtractString(c []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(c)
	if match != nil && len(match) >= 2 {
		return string(bytes.Join(match[1:], []byte{'_'}))
	} else {
		return ""
	}
}