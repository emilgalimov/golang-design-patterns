package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct_CreateProduct(t *testing.T) {
	product := NewProductLeaf()

	assert.IsType(t, &ProductLeaf{}, product)
	assert.NotEmpty(t, product.ID())
}

func TestProduct_UpdatesPrice(t *testing.T) {
	product := NewProductLeaf()
	product.UpdatePrice(123)
}

func TestProduct_ReturnPrice(t *testing.T) {
	product := NewProductLeaf()
	product.UpdatePrice(123)
	assert.Equal(t, 123, product.Price())
}

func TestProduct_IncreasesPricePercent(t *testing.T) {
	product := NewProductLeaf()
	product.UpdatePrice(100)
	product.IncreasePricePercent(5)
	assert.Equal(t, 105, product.Price())
}
