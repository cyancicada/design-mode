package main

import "design-mode/结构型/代理模式/proxy"

func main() {
	proxy := new(proxy.ProObject)
	proxy.ObjDo("Well")
}
