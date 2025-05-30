package decorator

import "fmt"

type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (sc *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (sc *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

type CoffeeDecorator struct {
	coffee Coffee
}

func (cd *CoffeeDecorator) Cost() float64 {
	return cd.coffee.Cost()
}

func (cd *CoffeeDecorator) Description() string {
	return cd.coffee.Description()
}

type MilkDecorator struct {
	CoffeeDecorator
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (md *MilkDecorator) Cost() float64 {
	return md.coffee.Cost() + 0.5
}

func (md *MilkDecorator) Description() string {
	return md.coffee.Description() + ", Milk"
}

type SugarDecorator struct {
	CoffeeDecorator
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (sd *SugarDecorator) Cost() float64 {
	return sd.coffee.Cost() + 0.2
}

func (sd *SugarDecorator) Description() string {
	return sd.coffee.Description() + ", Sugar"
}

type WhipDecorator struct {
	CoffeeDecorator
}

func NewWhipDecorator(coffee Coffee) *WhipDecorator {
	return &WhipDecorator{
		CoffeeDecorator: CoffeeDecorator{coffee: coffee},
	}
}

func (wd *WhipDecorator) Cost() float64 {
	return wd.coffee.Cost() + 0.7
}

func (wd *WhipDecorator) Description() string {
	return wd.coffee.Description() + ", Whip"
}

func PrintCoffeeInfo(coffee Coffee) {
	fmt.Printf("%s: $%.2f\n", coffee.Description(), coffee.Cost())
}