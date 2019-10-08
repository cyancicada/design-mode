package main

import "design-mode/行为型/观察者模式/observer"

func main(){
	eventCenter := observer.NewEventCenter()
	r_1 := observer.EventReciver{}
	r_2 := observer.EventReciver{}

	eventCenter.Register(&r_1)
	eventCenter.Register(&r_2)
	eventCenter.Notify(observer.Event{1})
	eventCenter.Degister(&r_1)
	eventCenter.Notify(observer.Event{2})
}