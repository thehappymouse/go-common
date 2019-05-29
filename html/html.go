package html

import (
	"golang.org/x/net/html"
)

// 获取一个元素节点的属性值
func GetNodeVal(key string, n *html.Node) (bool, string) {
	for _, a := range n.Attr {
		if a.Key == key {
			return true, a.Val
		}
	}
	return false, ""
}

// 设置一个元素节点的属性值
func SetNodeVal(key, val string, n *html.Node) bool {
	for index, a := range n.Attr {
		if a.Key == key {
			n.Attr[index].Val = val
			return true
		}
	}
	n.Attr = append(n.Attr, html.Attribute{Key: key, Val: val})
	return true
}

// 遍历每一个node，并执行回调
func VisitFn(n *html.Node, execute func(b *html.Node)) {
	execute(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		VisitFn(c, execute)
	}
}
