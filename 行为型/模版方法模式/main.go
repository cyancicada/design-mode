package main

import (
	"design-mode/行为型/模版方法模式/template"
)

func main() {
	a := template.TmplA{}
	b := template.TmplB{}
	template.Operate(&a)
	template.Operate(&b)

}
