package main

import "fmt"

func main() {

	employee := NewEmployee(15, 3)

	staffStatistic := NewEmployeeStatisticContext(
		employee,
		NewStaffSalaryStrategy(),
	)
	fmt.Printf("Employee salary as staff is %v\n", staffStatistic.CountSalary())

	freelancerStatistic := NewEmployeeStatisticContext(
		employee,
		NewFreelanceSalaryStrategy(),
	)
	fmt.Printf("Employee salary as freelancer is %v\n", freelancerStatistic.CountSalary())

	fmt.Printf("Average time per task is %v\n", freelancerStatistic.AverageTimePerTask())
}
