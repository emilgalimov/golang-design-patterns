package main

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfferComposite_Create(t *testing.T) {
	offer := NewOfferComposite()

	assert.IsType(t, &OfferComposite{}, offer)
	assert.NotEmpty(t, offer.ID())
}

func TestOfferComposite_AddsAndReturnsAndRemovesChildren(t *testing.T) {
	offer := NewOfferComposite()

	mc := minimock.NewController(t)
	component1 := NewComponentMock(mc).IDMock.Return(1)
	component2 := NewComponentMock(mc).IDMock.Return(2)

	offer.Add(component1)
	offer.Add(component2)

	assert.Equal(t, component1, offer.GetChild(component1.ID()))
	offer.Remove(component1)
	assert.Equal(t, nil, offer.GetChild(component1.ID()))

	assert.Equal(t, component2, offer.GetChild(component2.ID()))
}

func TestOfferComposite_Price(t *testing.T) {
	offer := NewOfferComposite()

	mc := minimock.NewController(t)

	component1 := NewComponentMock(mc).IDMock.Return(1)
	component1.PriceMock.Return(10)

	component2 := NewComponentMock(mc).IDMock.Return(2)
	component2.PriceMock.Return(20)

	offer.Add(component1)
	offer.Add(component2)

	assert.Equal(t, 30, offer.Price())
}

func TestOfferComposite_IncreasePricePercent(t *testing.T) {
	offer := NewOfferComposite()

	mc := minimock.NewController(t)

	component1 := NewComponentMock(mc).IDMock.Return(1)
	component1.IncreasePricePercentMock.Expect(8)

	component2 := NewComponentMock(mc).IDMock.Return(2)
	component2.IncreasePricePercentMock.Expect(8)

	offer.Add(component1)
	offer.Add(component2)

	offer.IncreasePricePercent(8)
}
