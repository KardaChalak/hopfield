package hopfield

import (
	"errors"
	"fmt"
	"strings"
)

// HopfieldNet represents the weights between the neurons
// Should be a symmetrical matrix
type HopfieldNet struct {
	Weights   [][]int
	NrNeurons int
}

func (h HopfieldNet) String() string {
	var sb strings.Builder
	for _, v := range h.Weights {
		sb.WriteString(fmt.Sprintf("%v\n", v))
	}
	return sb.String()
}

// Recall error corrects the argument with regard to
// the weights in the hopfield network and returns the
// error corrected pattern
func (h HopfieldNet) Recall(pattern []int) ([]int, error) {
	if h.NrNeurons != len(pattern) {
		return []int{}, errors.New("The pattern array has to have the same length as the number of neurons")
	}

	activation := make([]int, len(pattern))

	for epoch := 0; epoch < 1; epoch++ {

		for i := range pattern {
			sum := 0
			for j := range pattern {
				sum += h.Weights[i][j] * pattern[j]
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
func (h HopfieldNet) InsertPattern(pattern []int) error {
	if h.NrNeurons != len(pattern) {
		return errors.New("The pattern array has to have the same length as the number of neurons")
	}

	for i := 0; i < h.NrNeurons; i++ {
		for j := 0; j < h.NrNeurons; j++ {
			h.Weights[i][j] += pattern[i] * pattern[j]
			if i == j {
				// no self connection
				h.Weights[i][j] = 0
			}
		}
	}

	return nil
}

// NewNet creates a new Hopfield net and returns it.
func NewNet(nrNeurons int) (HopfieldNet, error) {
	if nrNeurons < 0 {
		return HopfieldNet{}, errors.New("nrNeurons cant be negative")
	}
	weights := make([][]int, nrNeurons)
	for i := range weights {
		weights[i] = make([]int, nrNeurons)
	}
	return HopfieldNet{weights, nrNeurons}, nil
}
