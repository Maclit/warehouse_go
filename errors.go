package main

import "fmt"

type ErrInput []string
type ErrColor string
type ErrMap string

func (e ErrInput) Error() string {
	return fmt.Sprintf("Can not accept more than 5 parameters and less than 3. number of parameters: %d", len(e))
}

func (e ErrMap) Error() string {
	return fmt.Sprintf("The input for the map is not a number or not greater than zero, string: %s ", string(e))
}

func (e ErrColor) Error() string {
	return fmt.Sprintf("This input does not correspond to an color accepted by the program, string: %s ", string(e))
}
