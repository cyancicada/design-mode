package main

import "design-mode/行为型/命令模式/command"

func main() {
	abcStock := command.NewStock()

	buyStockOrder := new(command.BuyStock)
	buyStockOrder.BuyStock(abcStock)

	sellStockOrder := new(command.SellStock)
	sellStockOrder.SellStock(abcStock)

	broker := new(command.Broker)
	broker.TakeOrder(buyStockOrder)
	broker.TakeOrder(sellStockOrder)

	broker.PlaceOrders()
}
