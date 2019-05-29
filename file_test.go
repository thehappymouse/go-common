package utils

import (
	"path"
	"testing"
)

func TestParseUrlBaseName(t *testing.T) {
	url := "http://imgsrc.baidu.com/forum/w%3D580/sign=1b239fd377cb0a4685228b315b62f63e/b20cf0246b600c339094a3e1164c510fd8f9a19d.jpg";
	right := "b20cf0246b600c339094a3e1164c510fd8f9a19d.jpg"
	res := path.Base(url)
	if (right != res) {
		t.Errorf("%s != %s", right, res)
	}
}
