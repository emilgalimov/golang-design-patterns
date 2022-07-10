package main

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/emilgalimov/golang-design-patterns/strategy.SalaryStrategy -o ./strategy/salary_strategy_mock_test.go -n SalaryStrategyMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// SalaryStrategyMock implements SalaryStrategy
type SalaryStrategyMock struct {
	t minimock.Tester

	funccountSalary          func(employee *Employee) (i1 int)
	inspectFunccountSalary   func(employee *Employee)
	aftercountSalaryCounter  uint64
	beforecountSalaryCounter uint64
	countSalaryMock          mSalaryStrategyMockcountSalary
}

// NewSalaryStrategyMock returns a mock for SalaryStrategy
func NewSalaryStrategyMock(t minimock.Tester) *SalaryStrategyMock {
	m := &SalaryStrategyMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.countSalaryMock = mSalaryStrategyMockcountSalary{mock: m}
	m.countSalaryMock.callArgs = []*SalaryStrategyMockcountSalaryParams{}

	return m
}

type mSalaryStrategyMockcountSalary struct {
	mock               *SalaryStrategyMock
	defaultExpectation *SalaryStrategyMockcountSalaryExpectation
	expectations       []*SalaryStrategyMockcountSalaryExpectation

	callArgs []*SalaryStrategyMockcountSalaryParams
	mutex    sync.RWMutex
}

// SalaryStrategyMockcountSalaryExpectation specifies expectation struct of the SalaryStrategy.CountSalary
type SalaryStrategyMockcountSalaryExpectation struct {
	mock    *SalaryStrategyMock
	params  *SalaryStrategyMockcountSalaryParams
	results *SalaryStrategyMockcountSalaryResults
	Counter uint64
}

// SalaryStrategyMockcountSalaryParams contains parameters of the SalaryStrategy.CountSalary
type SalaryStrategyMockcountSalaryParams struct {
	employee *Employee
}

// SalaryStrategyMockcountSalaryResults contains results of the SalaryStrategy.CountSalary
type SalaryStrategyMockcountSalaryResults struct {
	i1 int
}

// Expect sets up expected params for SalaryStrategy.CountSalary
func (mmcountSalary *mSalaryStrategyMockcountSalary) Expect(employee *Employee) *mSalaryStrategyMockcountSalary {
	if mmcountSalary.mock.funccountSalary != nil {
		mmcountSalary.mock.t.Fatalf("SalaryStrategyMock.CountSalary mock is already set by Set")
	}

	if mmcountSalary.defaultExpectation == nil {
		mmcountSalary.defaultExpectation = &SalaryStrategyMockcountSalaryExpectation{}
	}

	mmcountSalary.defaultExpectation.params = &SalaryStrategyMockcountSalaryParams{employee}
	for _, e := range mmcountSalary.expectations {
		if minimock.Equal(e.params, mmcountSalary.defaultExpectation.params) {
			mmcountSalary.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmcountSalary.defaultExpectation.params)
		}
	}

	return mmcountSalary
}

// Inspect accepts an inspector function that has same arguments as the SalaryStrategy.CountSalary
func (mmcountSalary *mSalaryStrategyMockcountSalary) Inspect(f func(employee *Employee)) *mSalaryStrategyMockcountSalary {
	if mmcountSalary.mock.inspectFunccountSalary != nil {
		mmcountSalary.mock.t.Fatalf("Inspect function is already set for SalaryStrategyMock.CountSalary")
	}

	mmcountSalary.mock.inspectFunccountSalary = f

	return mmcountSalary
}

// Return sets up results that will be returned by SalaryStrategy.CountSalary
func (mmcountSalary *mSalaryStrategyMockcountSalary) Return(i1 int) *SalaryStrategyMock {
	if mmcountSalary.mock.funccountSalary != nil {
		mmcountSalary.mock.t.Fatalf("SalaryStrategyMock.CountSalary mock is already set by Set")
	}

	if mmcountSalary.defaultExpectation == nil {
		mmcountSalary.defaultExpectation = &SalaryStrategyMockcountSalaryExpectation{mock: mmcountSalary.mock}
	}
	mmcountSalary.defaultExpectation.results = &SalaryStrategyMockcountSalaryResults{i1}
	return mmcountSalary.mock
}

//Set uses given function f to mock the SalaryStrategy.CountSalary method
func (mmcountSalary *mSalaryStrategyMockcountSalary) Set(f func(employee *Employee) (i1 int)) *SalaryStrategyMock {
	if mmcountSalary.defaultExpectation != nil {
		mmcountSalary.mock.t.Fatalf("Default expectation is already set for the SalaryStrategy.CountSalary method")
	}

	if len(mmcountSalary.expectations) > 0 {
		mmcountSalary.mock.t.Fatalf("Some expectations are already set for the SalaryStrategy.CountSalary method")
	}

	mmcountSalary.mock.funccountSalary = f
	return mmcountSalary.mock
}

// When sets expectation for the SalaryStrategy.CountSalary which will trigger the result defined by the following
// Then helper
func (mmcountSalary *mSalaryStrategyMockcountSalary) When(employee *Employee) *SalaryStrategyMockcountSalaryExpectation {
	if mmcountSalary.mock.funccountSalary != nil {
		mmcountSalary.mock.t.Fatalf("SalaryStrategyMock.CountSalary mock is already set by Set")
	}

	expectation := &SalaryStrategyMockcountSalaryExpectation{
		mock:   mmcountSalary.mock,
		params: &SalaryStrategyMockcountSalaryParams{employee},
	}
	mmcountSalary.expectations = append(mmcountSalary.expectations, expectation)
	return expectation
}

// Then sets up SalaryStrategy.CountSalary return parameters for the expectation previously defined by the When method
func (e *SalaryStrategyMockcountSalaryExpectation) Then(i1 int) *SalaryStrategyMock {
	e.results = &SalaryStrategyMockcountSalaryResults{i1}
	return e.mock
}

// countSalary implements SalaryStrategy
func (mmcountSalary *SalaryStrategyMock) countSalary(employee *Employee) (i1 int) {
	mm_atomic.AddUint64(&mmcountSalary.beforecountSalaryCounter, 1)
	defer mm_atomic.AddUint64(&mmcountSalary.aftercountSalaryCounter, 1)

	if mmcountSalary.inspectFunccountSalary != nil {
		mmcountSalary.inspectFunccountSalary(employee)
	}

	mm_params := &SalaryStrategyMockcountSalaryParams{employee}

	// Record call args
	mmcountSalary.countSalaryMock.mutex.Lock()
	mmcountSalary.countSalaryMock.callArgs = append(mmcountSalary.countSalaryMock.callArgs, mm_params)
	mmcountSalary.countSalaryMock.mutex.Unlock()

	for _, e := range mmcountSalary.countSalaryMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1
		}
	}

	if mmcountSalary.countSalaryMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmcountSalary.countSalaryMock.defaultExpectation.Counter, 1)
		mm_want := mmcountSalary.countSalaryMock.defaultExpectation.params
		mm_got := SalaryStrategyMockcountSalaryParams{employee}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmcountSalary.t.Errorf("SalaryStrategyMock.CountSalary got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmcountSalary.countSalaryMock.defaultExpectation.results
		if mm_results == nil {
			mmcountSalary.t.Fatal("No results are set for the SalaryStrategyMock.CountSalary")
		}
		return (*mm_results).i1
	}
	if mmcountSalary.funccountSalary != nil {
		return mmcountSalary.funccountSalary(employee)
	}
	mmcountSalary.t.Fatalf("Unexpected call to SalaryStrategyMock.CountSalary. %v", employee)
	return
}

// countSalaryAfterCounter returns a count of finished SalaryStrategyMock.countSalary invocations
func (mmcountSalary *SalaryStrategyMock) countSalaryAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmcountSalary.aftercountSalaryCounter)
}

// countSalaryBeforeCounter returns a count of SalaryStrategyMock.countSalary invocations
func (mmcountSalary *SalaryStrategyMock) countSalaryBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmcountSalary.beforecountSalaryCounter)
}

// Calls returns a list of arguments used in each call to SalaryStrategyMock.CountSalary.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmcountSalary *mSalaryStrategyMockcountSalary) Calls() []*SalaryStrategyMockcountSalaryParams {
	mmcountSalary.mutex.RLock()

	argCopy := make([]*SalaryStrategyMockcountSalaryParams, len(mmcountSalary.callArgs))
	copy(argCopy, mmcountSalary.callArgs)

	mmcountSalary.mutex.RUnlock()

	return argCopy
}

// MinimockcountSalaryDone returns true if the count of the CountSalary invocations corresponds
// the number of defined expectations
func (m *SalaryStrategyMock) MinimockcountSalaryDone() bool {
	for _, e := range m.countSalaryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.countSalaryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.aftercountSalaryCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funccountSalary != nil && mm_atomic.LoadUint64(&m.aftercountSalaryCounter) < 1 {
		return false
	}
	return true
}

// MinimockcountSalaryInspect logs each unmet expectation
func (m *SalaryStrategyMock) MinimockcountSalaryInspect() {
	for _, e := range m.countSalaryMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SalaryStrategyMock.CountSalary with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.countSalaryMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.aftercountSalaryCounter) < 1 {
		if m.countSalaryMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SalaryStrategyMock.CountSalary")
		} else {
			m.t.Errorf("Expected call to SalaryStrategyMock.CountSalary with params: %#v", *m.countSalaryMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funccountSalary != nil && mm_atomic.LoadUint64(&m.aftercountSalaryCounter) < 1 {
		m.t.Error("Expected call to SalaryStrategyMock.CountSalary")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SalaryStrategyMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockcountSalaryInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SalaryStrategyMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *SalaryStrategyMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockcountSalaryDone()
}