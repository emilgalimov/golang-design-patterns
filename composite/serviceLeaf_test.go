package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateService(t *testing.T) {
	service := NewServiceLeaf()
	assert.IsType(t, &ServiceLeaf{}, service)
	assert.NotEmpty(t, service.ID())
}

func TestService_UpdatesPrice(t *testing.T) {
	service := NewServiceLeaf()
	service.UpdatePrice(123)
}

func TestService_ReturnPrice(t *testing.T) {
	service := NewServiceLeaf()
	service.UpdatePrice(123)
	assert.Equal(t, 123, service.Price())
}

func TestService_IncreasesPricePercent(t *testing.T) {
	service := NewServiceLeaf()
	service.UpdatePrice(100)
	service.IncreasePricePercent(5)
	assert.Equal(t, 105, service.Price())
}
