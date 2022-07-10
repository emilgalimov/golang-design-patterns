package main

type FreelanceSalaryStrategy struct {
}

func NewFreelanceSalaryStrategy() *FreelanceSalaryStrategy {
	return &FreelanceSalaryStrategy{}
}

func (f *FreelanceSalaryStrategy) countSalary(employee *Employee) int {
	return employee.tasksCompleted * 1000
}
