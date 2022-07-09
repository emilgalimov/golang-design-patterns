package main

import "fmt"

func main() {

	washingMachine := NewProductLeaf()
	washingMachine.UpdatePrice(999)

	setUpWashingMachine := NewServiceLeaf()
	setUpWashingMachine.UpdatePrice(100)

	bathroom := NewOfferComposite()
	bathroom.Add(washingMachine)
	bathroom.Add(setUpWashingMachine)

	bricks := NewProductLeaf()
	bricks.UpdatePrice(15000)

	buildHouse := NewServiceLeaf()
	buildHouse.UpdatePrice(20000)

	home := NewOfferComposite()
	home.Add(bathroom)
	home.Add(bricks)
	home.Add(buildHouse)

	fmt.Printf("House price is %v\n", home.Price())

	home.IncreasePricePercent(10)

	fmt.Printf("New washing machine price is %v\n", washingMachine.Price())
	fmt.Printf("House price is %v\n", home.Price())
}
