package main

import "math/rand"

type OfferComposite struct {
	id       int
	children map[int]offerComponent
}

func NewOfferComposite() *OfferComposite {
	return &OfferComposite{
		id:       rand.Int(),
		children: make(map[int]offerComponent),
	}
}

func (o OfferComposite) ID() int {
	return o.id
}

func (o *OfferComposite) Add(child offerComponent) {
	o.children[child.ID()] = child
}

func (o OfferComposite) GetChild(id int) interface{} {
	return o.children[id]
}

func (o *OfferComposite) Remove(child offerComponent) {
	delete(o.children, child.ID())
}

func (o OfferComposite) Price() (p int) {
	for _, c := range o.children {
		p += c.Price()
	}
	return
}

func (o *OfferComposite) IncreasePricePercent(p int) {
	for _, c := range o.children {
		c.IncreasePricePercent(p)
	}
}
