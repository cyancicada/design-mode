package main

import (
	"fmt"
	"sync"
	"time"

	"design-mode/创建型/单例模式/singleton"
)

func main1() {
	instance_1 := singleton.GetInstance()

	instance_2 := singleton.GetInstance()
	fmt.Printf("instance_1=%p\ninstance_2=%p", instance_1, instance_2)
	//fmt.Println(s)
}

var once sync.Once

func main() {

	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
