package main

import . "design-mode/创建型/工厂模式/factories"

func main() {
	// Abstract Factory

	fa := FactoryA{}
	fa.CreateFood().Eat()
	fa.CreateDrink().Drink()

	fb := FactoryB{}
	fb.CreateFood().Eat()
	fb.CreateDrink().Drink()

}
