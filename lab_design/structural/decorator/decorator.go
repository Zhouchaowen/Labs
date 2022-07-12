package main

import "fmt"

type pizza interface {
	getPrice() int
}

type peppyPaneer struct {
}

func (p *peppyPaneer) getPrice() int {
	return 20
}

type veggeMania struct {
}

func (p *veggeMania) getPrice() int {
	return 15
}

type tomatoTopping struct { // 装饰原始pizza
	pizza pizza
}

func (c *tomatoTopping) getPrice() int { // 装饰原始pizza
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int { // 装饰原始pizza
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {

	veggiePizza := &veggeMania{}

	//Add cheese topping
	veggiePizzaWithCheese := &cheeseTopping{
		pizza: veggiePizza,
	}

	//Add tomato topping
	veggiePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.getPrice())

	peppyPaneerPizza := &peppyPaneer{}

	//Add cheese topping
	peppyPaneerPizzaWithCheese := &cheeseTopping{
		pizza: peppyPaneerPizza,
	}

	fmt.Printf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPaneerPizzaWithCheese.getPrice())
}
