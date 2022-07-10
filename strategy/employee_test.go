package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmployee_Create(t *testing.T) {
	employee := NewEmployee(
		30,
		4,
	)

	assert.Equal(t, 30, employee.WorkHours())
	assert.Equal(t, 4, employee.TasksCompleted())
}
