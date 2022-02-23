package main

import "fmt"

// InputError correspond to an error of Input
type InputError string

func (e InputError) Error() string {
	return fmt.Sprintf("Can not accept more than 5 parameters and less than 3. number of parameters: %s", string(e))
}

// MapError correspond to an error in the creation of the map
type MapError string

func (e MapError) Error() string {
	return fmt.Sprintf("The input for the map is not a number or not greater than zero, string: %s ", string(e))
}

// ColorError correspond to an error in the box color
type ColorError string

func (e ColorError) Error() string {
	return fmt.Sprintf("This input does not correspond to an color accepted by the program, string: %s ", string(e))
}

// StuckTransporterError correspond to when a transporter as no empty neighbor nodes
type StuckTransporterError string

func (e StuckTransporterError) Error() string {
	return fmt.Sprintf("Transporter is stuck. name : %s ", string(e))
}

// NotEmpytNodeError Returned when an empty node is not found.
type NotEmpytNodeError string

func (e NotEmpytNodeError) Error() string {
	return fmt.Sprintf("Node is not empty. name : %s ", string(e))
}

// NoNeighborNodeError Returned when no valid neigbor has been found for a node.
type NoNeighborNodeError string

func (e NoNeighborNodeError) Error() string {
	return fmt.Sprintf("Node has no neighbor. direction : %s ", string(e))
}
