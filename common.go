package utils

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"regexp"
	"strconv"
	"strings"
)

var p  = message.NewPrinter(language.English)

// 千分位展示
func GetMoney(v string) string {
	v2, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return v
	}
	return p.Sprintf("%.2f", v2)
}


// unicode 字符 转中文
func Unicode2Chinese(textUnquoted string) (string, error) {
	exp1 := regexp.MustCompile(`\\u([a-z0-9A-Z]{1,4})`)
	result1 := exp1.FindAllStringSubmatch(textUnquoted, -1)
	if len(result1) == 0 {
		return textUnquoted, nil
	}
	for _, s := range result1 {
		temp, err := strconv.ParseInt(s[1], 16, 32)
		if err != nil {
			return textUnquoted, err
		}
		textUnquoted = strings.Replace(textUnquoted, s[0], fmt.Sprintf("%c", temp), -1)
	}
	return textUnquoted, nil
}
