package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFreelanceSalaryStrategy_countSalary(t *testing.T) {
	strategy := NewFreelanceSalaryStrategy()

	employee := NewEmployee(10, 3)

	salary := strategy.countSalary(employee)

	assert.Equal(t, 3000, salary)
}
