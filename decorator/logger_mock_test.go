package main

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/emilgalimov/golang-design-patterns/decorator.requestLogger -o ./decorator/logger_mock_test.go -n LoggerMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// LoggerMock implements requestLogger
type LoggerMock struct {
	t minimock.Tester

	funcprint          func(s1 string)
	inspectFuncprint   func(s1 string)
	afterprintCounter  uint64
	beforeprintCounter uint64
	printMock          mLoggerMockprint
}

// NewLoggerMock returns a mock for requestLogger
func NewLoggerMock(t minimock.Tester) *LoggerMock {
	m := &LoggerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.printMock = mLoggerMockprint{mock: m}
	m.printMock.callArgs = []*LoggerMockprintParams{}

	return m
}

type mLoggerMockprint struct {
	mock               *LoggerMock
	defaultExpectation *LoggerMockprintExpectation
	expectations       []*LoggerMockprintExpectation

	callArgs []*LoggerMockprintParams
	mutex    sync.RWMutex
}

// LoggerMockprintExpectation specifies expectation struct of the requestLogger.print
type LoggerMockprintExpectation struct {
	mock   *LoggerMock
	params *LoggerMockprintParams

	Counter uint64
}

// LoggerMockprintParams contains parameters of the requestLogger.print
type LoggerMockprintParams struct {
	s1 string
}

// Expect sets up expected params for requestLogger.print
func (mmprint *mLoggerMockprint) Expect(s1 string) *mLoggerMockprint {
	if mmprint.mock.funcprint != nil {
		mmprint.mock.t.Fatalf("LoggerMock.print mock is already set by Set")
	}

	if mmprint.defaultExpectation == nil {
		mmprint.defaultExpectation = &LoggerMockprintExpectation{}
	}

	mmprint.defaultExpectation.params = &LoggerMockprintParams{s1}
	for _, e := range mmprint.expectations {
		if minimock.Equal(e.params, mmprint.defaultExpectation.params) {
			mmprint.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmprint.defaultExpectation.params)
		}
	}

	return mmprint
}

// Inspect accepts an inspector function that has same arguments as the requestLogger.print
func (mmprint *mLoggerMockprint) Inspect(f func(s1 string)) *mLoggerMockprint {
	if mmprint.mock.inspectFuncprint != nil {
		mmprint.mock.t.Fatalf("Inspect function is already set for LoggerMock.print")
	}

	mmprint.mock.inspectFuncprint = f

	return mmprint
}

// Return sets up results that will be returned by requestLogger.print
func (mmprint *mLoggerMockprint) Return() *LoggerMock {
	if mmprint.mock.funcprint != nil {
		mmprint.mock.t.Fatalf("LoggerMock.print mock is already set by Set")
	}

	if mmprint.defaultExpectation == nil {
		mmprint.defaultExpectation = &LoggerMockprintExpectation{mock: mmprint.mock}
	}

	return mmprint.mock
}

//Set uses given function f to mock the requestLogger.print method
func (mmprint *mLoggerMockprint) Set(f func(s1 string)) *LoggerMock {
	if mmprint.defaultExpectation != nil {
		mmprint.mock.t.Fatalf("Default expectation is already set for the requestLogger.print method")
	}

	if len(mmprint.expectations) > 0 {
		mmprint.mock.t.Fatalf("Some expectations are already set for the requestLogger.print method")
	}

	mmprint.mock.funcprint = f
	return mmprint.mock
}

// print implements requestLogger
func (mmprint *LoggerMock) print(s1 string) {
	mm_atomic.AddUint64(&mmprint.beforeprintCounter, 1)
	defer mm_atomic.AddUint64(&mmprint.afterprintCounter, 1)

	if mmprint.inspectFuncprint != nil {
		mmprint.inspectFuncprint(s1)
	}

	mm_params := &LoggerMockprintParams{s1}

	// Record call args
	mmprint.printMock.mutex.Lock()
	mmprint.printMock.callArgs = append(mmprint.printMock.callArgs, mm_params)
	mmprint.printMock.mutex.Unlock()

	for _, e := range mmprint.printMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmprint.printMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmprint.printMock.defaultExpectation.Counter, 1)
		mm_want := mmprint.printMock.defaultExpectation.params
		mm_got := LoggerMockprintParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmprint.t.Errorf("LoggerMock.print got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmprint.funcprint != nil {
		mmprint.funcprint(s1)
		return
	}
	mmprint.t.Fatalf("Unexpected call to LoggerMock.print. %v", s1)

}

// printAfterCounter returns a count of finished LoggerMock.print invocations
func (mmprint *LoggerMock) printAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmprint.afterprintCounter)
}

// printBeforeCounter returns a count of LoggerMock.print invocations
func (mmprint *LoggerMock) printBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmprint.beforeprintCounter)
}

// Calls returns a list of arguments used in each call to LoggerMock.print.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmprint *mLoggerMockprint) Calls() []*LoggerMockprintParams {
	mmprint.mutex.RLock()

	argCopy := make([]*LoggerMockprintParams, len(mmprint.callArgs))
	copy(argCopy, mmprint.callArgs)

	mmprint.mutex.RUnlock()

	return argCopy
}

// MinimockprintDone returns true if the count of the print invocations corresponds
// the number of defined expectations
func (m *LoggerMock) MinimockprintDone() bool {
	for _, e := range m.printMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.printMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterprintCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcprint != nil && mm_atomic.LoadUint64(&m.afterprintCounter) < 1 {
		return false
	}
	return true
}

// MinimockprintInspect logs each unmet expectation
func (m *LoggerMock) MinimockprintInspect() {
	for _, e := range m.printMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LoggerMock.print with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.printMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterprintCounter) < 1 {
		if m.printMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to LoggerMock.print")
		} else {
			m.t.Errorf("Expected call to LoggerMock.print with params: %#v", *m.printMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcprint != nil && mm_atomic.LoadUint64(&m.afterprintCounter) < 1 {
		m.t.Error("Expected call to LoggerMock.print")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *LoggerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockprintInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *LoggerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *LoggerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockprintDone()
}