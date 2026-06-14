package main

import (
	"awesomeProject3/Client"
	"errors"
	"fmt"
	"sync"
)

type ServerError struct {
	err error
}

func (s *ServerError) Error() string {
	return s.err.Error()
}

// between
func newMessageError(msg string) error {
	return &ServerError{err: errors.New(msg)}
}

var internetError = newMessageError("Internal system error")

func getClient() (Client.Client, error) {
	value := Client.Client{}
	err := fmt.Errorf("getClient: client received error from our system : %w", internetError)
	return value, err
}

func divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("a and b cannot be 0")
	}

	return a / b, nil
}

func main() {

	var f_value float64
	var s_value float64

	fmt.Print("Enter first value : ")
	fmt.Scan(&f_value)

	fmt.Print("Enter second value: ")
	fmt.Scan(&s_value)

	var wg sync.WaitGroup
	var divRes float64
	var divErr error
	var cli Client.Client
	var cliErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		divRes, divErr = divide(f_value, s_value)
	}()

	go func() {
		defer wg.Done()
		cli, cliErr = getClient()
	}()

	wg.Wait()

	if divErr != nil {
		fmt.Println("Divide error : ", divErr)
		return
	}
	fmt.Println("Divide result : ", divRes)

	if cliErr != nil {
		if errors.Is(cliErr, internetError) {
			fmt.Println("An internal error has occurred")
		}
		fmt.Println("Raw error : ", errors.Unwrap(cliErr))
		return
	}

	fmt.Println("Client : ", cli)
}
