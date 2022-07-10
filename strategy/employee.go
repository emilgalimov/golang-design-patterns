package main

type Employee struct {
	workHours      int
	tasksCompleted int
}

func NewEmployee(workHours int, tasksCompleted int) *Employee {
	return &Employee{
		workHours:      workHours,
		tasksCompleted: tasksCompleted,
	}
}

func (e *Employee) WorkHours() int {
	return e.workHours
}

func (e *Employee) TasksCompleted() int {
	return e.tasksCompleted
}
