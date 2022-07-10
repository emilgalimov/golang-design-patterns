package main

type SalaryStrategy interface {
	countSalary(employee *Employee) int
}

type EmployeeStatisticContext struct {
	employee       *Employee
	salaryStrategy SalaryStrategy
}

func NewEmployeeStatisticContext(
	employee *Employee,
	strategy SalaryStrategy,
) *EmployeeStatisticContext {
	return &EmployeeStatisticContext{
		employee:       employee,
		salaryStrategy: strategy,
	}
}

func (c *EmployeeStatisticContext) Employee() interface{} {
	return c.employee
}

func (c *EmployeeStatisticContext) CountSalary() int {
	return c.salaryStrategy.countSalary(c.employee)
}

func (c *EmployeeStatisticContext) AverageTimePerTask() float64 {
	return float64(c.employee.workHours) / float64(c.employee.tasksCompleted)
}
