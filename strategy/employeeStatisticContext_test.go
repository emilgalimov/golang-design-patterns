package main

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmployeeStatisticContext_Create(t *testing.T) {

	employee := NewEmployee(10, 3)

	mc := minimock.NewController(t)
	defer mc.Finish()
	strategy := NewSalaryStrategyMock(mc)

	statisticContext := NewEmployeeStatisticContext(
		employee,
		strategy,
	)

	assert.IsType(t, &EmployeeStatisticContext{}, statisticContext)
	assert.Equal(t, employee, statisticContext.Employee())
}

func TestEmployeeStatisticContext_CallsSalaryStrategy(t *testing.T) {

	employee := NewEmployee(10, 3)

	mc := minimock.NewController(t)
	defer mc.Finish()
	strategy := NewSalaryStrategyMock(mc).countSalaryMock.Expect(employee).Return(123)

	statisticContext := NewEmployeeStatisticContext(
		employee,
		strategy,
	)

	assert.Equal(t, 123, statisticContext.CountSalary())
}

func TestEmployeeStatisticContext_AverageTimePerTask(t *testing.T) {
	employee := NewEmployee(15, 3)

	mc := minimock.NewController(t)
	defer mc.Finish()
	strategy := NewSalaryStrategyMock(mc)

	statisticContext := NewEmployeeStatisticContext(
		employee,
		strategy,
	)

	assert.Equal(t, 5.0, statisticContext.AverageTimePerTask())
}
