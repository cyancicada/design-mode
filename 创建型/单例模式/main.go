package main模式

import (
	"fmt"

	"design-mode/创建型/单例模式/singleton"
)

func main() {
	instance_1 := singleton.GetInstance()
	instance_1["this"] = "that"

	instance_2 := singleton.GetInstance()
	s := instance_2["this"]
	fmt.Println(s)
}
