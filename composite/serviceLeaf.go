package main

import "math/rand"

type ServiceLeaf struct {
	id    int
	price int
}

func (s ServiceLeaf) ID() int {
	return s.id
}

func (s ServiceLeaf) Price() int {
	return s.price
}

func (s *ServiceLeaf) UpdatePrice(p int) {
	s.price = p
}

func (s *ServiceLeaf) IncreasePricePercent(percent int) {
	s.price += s.price * percent / 100
}

func NewServiceLeaf() *ServiceLeaf {
	return &ServiceLeaf{
		id: rand.Int(),
	}
}
