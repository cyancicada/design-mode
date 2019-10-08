package main

import (
	"fmt"
	"time"

	"design-mode/行为型/命令模式/command"
)

func main() {
	sys := command.NewEventSystem()

	sys.Map(command.EVENT_CODE_KEY, func(data command.Event) {
		fmt.Println(data)
	})

	sys.InspectKeyboard()
	time.Sleep(time.Second * 100)
}
