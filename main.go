package main

import (
	"errors"
	"fmt"
)

type DomainErrorType int
type State int

const (
	NetworkRequest DomainErrorType = iota
	JsonParsing
	Storage
)

type DomainError struct {
	Type  DomainErrorType
	Error error
}

const (
	Success State = iota
	Failure
)

type Result[Output any, Fault any] struct {
	state State
	fault Fault
	value Output
}

func (result *Result[Output, Fault]) SetValue(value Output) {
	result.value = value
	result.state = Success
}

func (result *Result[Output, Fault]) Value() Output {
	return result.value
}

func (result *Result[Output, Fault]) SetFault(failure Fault) {
	result.fault = failure
	result.state = Failure
}

func (result *Result[Output, Fault]) Fault() Fault {
	return result.fault
}

func (result *Result[Output, Fault]) State() State {
	return result.state
}

func testSuccessOperation() *Result[string, DomainError] {
	result := new(Result[string, DomainError])
	result.SetValue("hi :) every things work fine.")

	return result
}

func testFailureOperation() *Result[string, DomainError] {
	result := new(Result[string, DomainError])

	myError := DomainError{
		Type:  Storage,
		Error: errors.New("database service unavailable"),
	}

	result.SetFault(myError)

	return result
}

func testFailureWithErrorType() *Result[string, error] {
	result := new(Result[string, error])

	goError := errors.New("this is an error using the Go `error` type")

	result.SetFault(goError)

	return result
}

func main() {
	testOk := testSuccessOperation()

	switch testOk.State() {
	case Success:
		fmt.Printf("👍: %s\n", testOk.Value())
	case Failure:
		fmt.Printf("🚨: %v\n", testOk.Fault())
	}

	testFailure := testFailureOperation()

	switch testFailure.State() {
	case Success:
		fmt.Printf("👍: %s\n", testFailure.Value())
	case Failure:
		fmt.Printf("🚨: %v\n", testFailure.Fault().Error)
	}

	testGoError := testFailureWithErrorType()

	switch testGoError.State() {
	case Success:
		fmt.Printf("👍: %s\n", testGoError.Value())
	case Failure:
		fmt.Printf("🚨: %v\n", testGoError.Fault())
	}
}
