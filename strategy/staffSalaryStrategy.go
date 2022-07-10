package main

type StaffSalaryStrategy struct {
}

func NewStaffSalaryStrategy() *StaffSalaryStrategy {
	return &StaffSalaryStrategy{}
}

func (s *StaffSalaryStrategy) countSalary(employee *Employee) int {
	return employee.workHours * 50
}
