
package main

import "fmt"

// Component interface
type Coffee interface {
	Cost() int
	Description() string
}

// Concrete component
type SimpleCoffee struct{}

func (c SimpleCoffee) Cost() int {
	return 50
}

func (c SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator: adds milk
type MilkDecorator struct {
	coffee Coffee
}

func (m MilkDecorator) Cost() int {
	return m.coffee.Cost() + 10
}

func (m MilkDecorator) Description() string {
	return m.coffee.Description() + " + Milk"
}

// Decorator: adds sugar
type SugarDecorator struct {
	coffee Coffee
}

func (s SugarDecorator) Cost() int {
	return s.coffee.Cost() + 5
}

func (s SugarDecorator) Description() string {
	return s.coffee.Description() + " + Sugar"
}

func main() {
	var coffee Coffee = SimpleCoffee{}

	// Dynamically decorate the coffee
	coffee = MilkDecorator{coffee: coffee}
	coffee = SugarDecorator{coffee: coffee}

	fmt.Println("Order:", coffee.Description())
	fmt.Println("Price: ₹", coffee.Cost())
}
