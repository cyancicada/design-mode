package command

import "fmt"

type Broker struct {
	orders []Order
}

func (b *Broker) TakeOrder(order Order) {
	b.orders = append(b.orders, order)
}

func (b *Broker) PlaceOrders() {
	for _, order := range b.orders {
		order.Execute()
	}
}

type Order interface {
	Execute()
}

type BuyStock struct {
	abcStock *Stock
}

func (b *BuyStock) BuyStock(abcStock *Stock) {
	b.abcStock = abcStock
}

func (b *BuyStock) Execute() {
	b.abcStock.Buy()
}

type SellStock struct {
	abcStock *Stock
}

func (s *SellStock) SellStock(abcStock *Stock) {
	s.abcStock = abcStock
}

func (s *SellStock) Execute() {
	s.abcStock.Sell()
}

type Stock struct {
	Name     string
	Quantity int
}

func NewStock() *Stock {
	stock := new(Stock)
	stock.Name = "ABC"
	stock.Quantity = 10
	return stock
}

func (s *Stock) Buy() {
	fmt.Println("Stock [ Name: ", s.Name, ", Quantity: ", s.Quantity, " ] bought")
}

func (s *Stock) Sell() {
	fmt.Println("Stock [ Name: ", s.Name, ",Quantity: ", s.Quantity, " ] sold")
}
