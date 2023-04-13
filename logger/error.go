package logger

import (
	"fmt"
	"time"
)

// Implement a step based logger, steps are on the same line
type StepLogger struct {
	// The current step
	step int
	// The time at which the step started
	start time.Time
	// Errored
	err bool
	// Max step
	max int
}

func NewStepLogger(max int) *StepLogger {
	return &StepLogger{
		step:  1,
		start: time.Now(),
		max:   max,
	}
}

// Increment the step
func (l *StepLogger) Step() {
	fmt.Printf(" \033[1;32m(%s)\033[0m", time.Since(l.start))
	l.step++
	l.start = time.Now()
	fmt.Printf("\n")
}

func (l *StepLogger) StepString() string {
	return fmt.Sprintf("\033[1;34mStep %d/%d\033[0m", l.step, l.max)
}

func (l *StepLogger) Log(msg string) {
	// Empty the current line
	fmt.Printf("\r\033[K")
	fmt.Printf("\r%s: %s", l.StepString(), msg)
}

func (l *StepLogger) Logf(format string, args ...interface{}) {
	// Empty the current line
	fmt.Printf("\r\033[K")
	fmt.Printf("\r%s: %s", l.StepString(), fmt.Sprintf(format, args...))
}

func (l *StepLogger) LogError(msg string) {
	fmt.Printf("\r\033[K")
	fmt.Printf("\r\033[1;31m%s: %s\033[0m", l.StepString(), msg)
}

func (l *StepLogger) LogErrorf(format string, args ...interface{}) {
	fmt.Printf("\r\033[K")
	fmt.Printf("\r\033[1;31m%s: %s\033[0m", l.StepString(), fmt.Sprintf(format, args...))
}

func (l *StepLogger) LogErrorAfter(msg string) {
	if !l.err {
		l.LogError("Failed")
		l.err = true
	}
	fmt.Printf("\n")
	fmt.Printf("\033[1;31m%s\033[0m\n", msg)
}

func (l *StepLogger) LogErrorAfterf(format string, args ...interface{}) {
	if !l.err {
		l.LogError("Failed")
		l.err = true
	}
	fmt.Printf("\n")
	fmt.Printf("\033[1;31m%s\033[0m\n", fmt.Sprintf(format, args...))
}
