package main

import "math/rand"

type ProductLeaf struct {
	id    int
	price int
}

func NewProductLeaf() *ProductLeaf {
	return &ProductLeaf{
		id: rand.Int(),
	}
}

func (p ProductLeaf) ID() int {
	return p.id
}

func (p ProductLeaf) Price() int {
	return p.price
}

func (p *ProductLeaf) UpdatePrice(price int) {
	p.price = price
}

func (p *ProductLeaf) IncreasePricePercent(percent int) {
	p.price += p.price * percent / 100
}
