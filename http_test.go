package utils

import (
	"testing"
)

func TestHttpGetFile(t *testing.T) {
	url := "https://www.baidu.com/img/bd_logo1.png"
	local := "baidu.png"
	err := HttpGetFile(url, local)
	CheckError(err)

	if !IsFileExists(local) {
		t.Errorf("文件下载未成功:%s 不存在",  local)
	}
}
