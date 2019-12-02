package utils

import "testing"

func TestUnicode2Chinese(t *testing.T) {
	testData := []struct {
		Src   string
		Right string
	}{
		{Src: "ds", Right: "\u0064\u0073"},
		{Src: "国", Right: "\u56fd"},
		{Src: "ABCD\u8fd9\u91cc\u662f{}\u5317\u4eac}{\u70ed\u6cea}\u56fd", Right: "ABCD这里是{}北京}{热泪}国"},
	}
	for _, data := range testData {
		str, err := Unicode2Chinese(data.Src)
		if err != nil {
			t.Error(err)
		}
		if str != data.Right {
			t.Errorf("Unicode2Chinese结果：%s,正确值应该是:%s", str, data.Right)
		}
	}
}
