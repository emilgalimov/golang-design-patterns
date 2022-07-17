package main

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogRequestDecorator_Create(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	baseService := NewAppServiceMock(mc)
	loggerMock := NewLoggerMock(mc)

	log := NewLogRequestDecorator(baseService, loggerMock)

	assert.IsType(t, &logRequestDecorator{}, log)
}

func TestLogRequestDecorator_ExecuteBaseFunctionAndWritesLog(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	request := "request string"
	response := "response string"

	baseService := NewAppServiceMock(mc).processRequestMock.Expect(request).Return(response)
	loggerMock := NewLoggerMock(mc).printMock.Expect(request).Return()

	loggerDecorator := NewLogRequestDecorator(baseService, loggerMock)

	decoratorResponse := loggerDecorator.processRequest(request)

	assert.IsType(t, response, decoratorResponse)
}
