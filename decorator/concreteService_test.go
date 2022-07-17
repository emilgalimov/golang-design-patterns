package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcreteService_ItCanBeCreated(t *testing.T) {
	service := NewConcreteService()

	assert.IsType(t, &concreteService{}, service)
}

func TestConcreteService_ItProcessRequest(t *testing.T) {
	service := NewConcreteService()

	response := service.processRequest("message")
	assert.Equal(t, "Response is <message>", response)

	response2 := service.processRequest("new message")
	assert.Equal(t, "Response is <new message>", response2)
}
