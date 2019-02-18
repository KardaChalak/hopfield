package hopfield

import (
	"errors"
	"fmt"
	"strings"
)

// Net represents the weights between the neurons
// Should be a symmetrical matrix
type Net struct {
	Weights   [][]int
	NrNeurons int
}

func (net Net) String() string {
	var sb strings.Builder
	for _, v := range net.Weights {
		sb.WriteString(fmt.Sprintf("%v\n", v))
	}
	return sb.String()
}

// Recall error corrects the argument with regard to
// the weights in the hopfield network and returns the
// error corrected pattern
func (net Net) Recall(pattern []int) ([]int, error) {
	if net.NrNeurons != len(pattern) {
		return []int{}, errors.New("The pattern array has to have the same length as the number of neurons")
	}

	activation := make([]int, len(pattern))

	for epoch := 0; epoch < 1; epoch++ {

		for i := range pattern {
			sum := 0
			for j := range pattern {
				sum += net.Weights[i][j] * pattern[j]
			}
			if sum >= 0 {
				activation[i] = 1
			} else {
				activation[i] = -1
			}
		}
	}
	return activation, nil
}

// InsertPattern inserts a new pattern to the hopfield network
func (net Net) InsertPattern(pattern []int) error {
	if net.NrNeurons != len(pattern) {
		return errors.New("The pattern array has to have the same length as the number of neurons")
	}

	for i := 0; i < net.NrNeurons; i++ {
		for j := 0; j < net.NrNeurons; j++ {
			net.Weights[i][j] += pattern[i] * pattern[j]
			if i == j {
				// no self connection
				net.Weights[i][j] = 0
			}
		}
	}

	return nil
}

// NewNet creates a new Hopfield net and returns it.
func NewNet(nrNeurons int) (Net, error) {
	if nrNeurons < 0 {
		return Net{}, errors.New("nrNeurons cant be negative")
	}
	weights := make([][]int, nrNeurons)
	for i := range weights {
		weights[i] = make([]int, nrNeurons)
	}
	return Net{weights, nrNeurons}, nil
}
