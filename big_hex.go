package utils

import (
	"math"
	"strings"
	"time"
)

// 十进制转36进制
var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"

// 按 yymmdd-[A-z][0-Z] 形式生成序号
func UssNoGenerate(last string) string {
	prefix := time.Now().Format(DateTimeDate)
	value := 0
	if last != "" {
		suffix := strings.Split(last, "-")[1]
		value = BHex2Num(suffix, 36) + 1
	} else {
		value = BHex2Num("A0", 36)
	}
	return prefix + "-" + NumToBHex(value, 36)
}


// 10进制数转换   n 表示进制， 16 or 36
func NumToBHex(num, n int) string {
	num_str := ""
	for num != 0 {
		yu := num % n
		num_str = string(num2char[yu]) + num_str
		num = num / n
	}
	return strings.ToUpper(num_str)
}

// 36进制数转换   n 表示进制， 16 or 36
func BHex2Num(str string, n int) int {
	str = strings.ToLower(str)
	v := 0.0
	length := len(str)
	for i := 0; i < length; i++ {
		s := string(str[i])
		index := strings.Index(num2char, s)
		v += float64(index) * math.Pow(float64(n), float64(length-1-i)) // 倒序
	}
	return int(v)
}
