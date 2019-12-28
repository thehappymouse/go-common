package utils

import (
	"os"
	"testing"
)

func TestHttpGetFile(t *testing.T) {
	local := "baidu.png"
	os.Remove(local)
	url := "https://www.baidu.com/img/bd_logo1.png"
	HttpGetFile(url, local)
	if !IsFileExists(local) {
		t.Errorf("文件下载未成功:%s 不存在", local)
	}
	os.Remove(local)
}

func TestHttpGetFileFail(t *testing.T) {
	local := "baidu.png"
	os.Remove(local)
	url := "https://ast-oss.oss-cn-beijing.aliyuncs.com/order-pdf"
	err := HttpGetFile(url, local)
	if err == nil {
		t.Errorf("%s 不存在的地址，下载不可能成功", url)
	}
}
