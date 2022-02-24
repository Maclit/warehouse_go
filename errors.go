package main

import "fmt"

// ArgumentsError correspond to an error of executable arguments
type ArgumentsError string

func (e ArgumentsError) Error() string {
	return fmt.Sprintf("Bad parameters: %s", string(e))
}

// InputError correspond to an error of Input
type InputError string

func (e InputError) Error() string {
	return fmt.Sprintf("Can not accept more than 5 parameters and less than 3. number of parameters: %s", string(e))
}

// MapError correspond to an error in the creation of the map
type MapError string

func (e MapError) Error() string {
	return fmt.Sprintf("The input is invalid, string: %s ", string(e))
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

// GraphError Returned when graph content are not valid.
type GraphError string

func (e GraphError) Error() string {
	return fmt.Sprintf("Invalid graph content : %s ", string(e))
}

// BadGraphCoordinatesError Returned when coordinates passed to the graph does not exist.
type BadGraphCoordinatesError string

func (e BadGraphCoordinatesError) Error() string {
	return fmt.Sprintf("Invalid graph coordinates : %s ", string(e))
}
