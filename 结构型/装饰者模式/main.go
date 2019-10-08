package main

import (
	"fmt"
	"math"

	"design-mode/结构型/装饰者模式/decorator"
)

func Double(decoratorFunc decorator.DecoratorFunc) decorator.DecoratorFunc {
	return func(f float64) float64 {
		var result float64 = f
		if decoratorFunc != nil {
			result = decoratorFunc(f)
		}
		return result * 2
	}
}

func Sqrt(decoratorFunc decorator.DecoratorFunc) decorator.DecoratorFunc {
	return func(f float64) float64 {
		var result float64 = f
		if decoratorFunc != nil {
			result = decoratorFunc(f)
		}
		return math.Sqrt(result)
	}
}

func main() {
	decorator.Log("Yeah", "WWT")

	f := decorator.DecFunc(Double(Sqrt(nil)))
	fmt.Println(f(16.0))
}
