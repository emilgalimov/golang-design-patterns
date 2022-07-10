package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaffSalaryStrategy_countSalary(t *testing.T) {
	strategy := NewStaffSalaryStrategy()

	employee := NewEmployee(100, 3)

	salary := strategy.countSalary(employee)

	assert.Equal(t, 5000, salary)
}
