package main

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

	for epoch := 0; epoch < 10; epoch++ {

		for i := range pattern {
			sum := 0
			for j := range pattern {
				sum += h.Weights[i][j] * activation[j]
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
// and return a true boolean if it was successfull
func (h HopfieldNet) InsertPattern(pattern []int) (bool, error) {
	if h.NrNeurons != len(pattern) {
		return false, errors.New("The pattern array has to have the same length as the number of neurons")
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

	return true, nil
}

// NewNet creates a new Hopfield net and returns it.
func NewNet(NrNeurons int) HopfieldNet {
	weights := make([][]int, NrNeurons)
	for i := range weights {
		weights[i] = make([]int, NrNeurons)
	}
	return HopfieldNet{weights, NrNeurons}
}

func main() {
	train1 := []int{1, 1, 1, 1, 1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1}
	train2 := []int{1, 1, 1, 1, 1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1}
	train3 := []int{1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1}
	test := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, 1}
	hNet := NewNet(24)
	hNet.InsertPattern(train1)
	hNet.InsertPattern(train2)
	hNet.InsertPattern(train3)

	fmt.Println(train1)
	fmt.Println(hNet.Recall(test))
}
