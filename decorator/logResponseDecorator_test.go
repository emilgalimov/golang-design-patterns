package main

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogResponseDecorator_Create(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	baseService := NewAppServiceMock(mc)
	loggerMock := NewLoggerMock(mc)

	log := NewLogResponseDecorator(baseService, loggerMock)

	assert.IsType(t, &logResponseDecorator{}, log)
}

func TestLogResponseDecorator_ExecuteBaseFunctionAndWritesLog(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	request := "request string"
	response := "response string"

	baseService := NewAppServiceMock(mc).processRequestMock.Expect(request).Return(response)
	loggerMock := NewLoggerMock(mc).printMock.Expect(response).Return()

	loggerDecorator := NewLogResponseDecorator(baseService, loggerMock)

	decoratorResponse := loggerDecorator.processRequest(request)

	assert.IsType(t, response, decoratorResponse)
}
