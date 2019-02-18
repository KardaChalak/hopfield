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

func (n Net) String() string {
	var sb strings.Builder
	for _, v := range n.Weights {
		sb.WriteString(fmt.Sprintf("%v\n", v))
	}
	return sb.String()
}

// Recall error corrects the argument with regard to
// the weights in the hopfield network and returns the
// error corrected pattern
func (n Net) Recall(pattern []int) ([]int, error) {
	if n.NrNeurons != len(pattern) {
		return []int{}, errors.New("The pattern array has to have the same length as the number of neurons")
	}

	activation := make([]int, len(pattern))

	for epoch := 0; epoch < 1; epoch++ {

		for i := range pattern {
			sum := 0
			for j := range pattern {
				sum += n.Weights[i][j] * pattern[j]
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
func (n Net) InsertPattern(p []int) error {
	if n.NrNeurons != len(p) {
		return errors.New("The pattern array has to have the same length as the number of neurons")
	}

	for i := 0; i < n.NrNeurons; i++ {
		for j := 0; j < n.NrNeurons; j++ {
			n.Weights[i][j] += p[i] * p[j]
			if i == j {
				// no self connection
				n.Weights[i][j] = 0
			}
		}
	}

	return nil
}

// NewNet creates a new Hopfield Network and returns it.
func NewNet(nrNeurons int) (Net, error) {
	if nrNeurons < 0 {
		return Net{}, errors.New("nrNeurons can't be negative")
	}
	weights := make([][]int, nrNeurons)
	for i := range weights {
		weights[i] = make([]int, nrNeurons)
	}
	return Net{weights, nrNeurons}, nil
}
