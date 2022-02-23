package main

import "fmt"

// ErrInput correspond to an error of Input
type ErrInput []string

func (e ErrInput) Error() string {
	return fmt.Sprintf("Can not accept more than 5 parameters and less than 3. number of parameters: %d", len(e))
}

// ErrMap correspond to an error in the creation of the map
type ErrMap string

func (e ErrMap) Error() string {
	return fmt.Sprintf("The input for the map is not a number or not greater than zero, string: %s ", string(e))
}

// ErrColor correspond to an error in the box color
type ErrColor string

func (e ErrColor) Error() string {
	return fmt.Sprintf("This input does not correspond to an color accepted by the program, string: %s ", string(e))
}

// ErrEmptyNeighbor correspond to when a transporter as no empty neighbor nodes
type ErrStuck string

func (e ErrStuck) Error() string {
	return fmt.Sprintf("Transporter is stuck. name : %s ", string(e))
}
