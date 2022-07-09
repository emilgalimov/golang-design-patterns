package main

type offerComponent interface {
	ID() int
	Price() int
	IncreasePricePercent(int)
}
