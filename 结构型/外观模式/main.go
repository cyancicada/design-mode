package main

import (
	"fmt"

	. "design-mode/结构型/外观模式/facade"
)

func EatVegtables(veg ...*Vegetable) {
	for _, v := range veg {
		Eat(v)
	}
}

func main() {

	// No Facade
	bc := BuyVegetable("白菜")
	SaveVegetables(bc)
	vecs := CookVegtable(GetVegetables())
	EatVegtables(vecs...)
	for _, v := range vecs {
		fmt.Println(*v)
	}

	// Favade
	sauteVegtable := SauteVegtable()
	EatVegtables(sauteVegtable...)
	for _, v := range sauteVegtable {
		fmt.Println(*v)
	}
}
