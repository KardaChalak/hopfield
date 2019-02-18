package main

import (
	"fmt"
	"strings"
)

// HopfieldNet represents the weights between the neurons
// Should be a symmetrical matrix
type HopfieldNet struct {
	Weights   [][]float64
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
func (h HopfieldNet) Recall(pattern []float64) []float64 {

	activation := make([]float64, len(pattern))

	for epoch := 0; epoch < 10; epoch++ {

		for i := range pattern {
			sum := 0.0
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
	return activation
}

// InsertPattern inserts a new pattern to the hopfield network
// and return a true boolean if it was successfull
func (h HopfieldNet) InsertPattern(pattern []float64) bool {

	for i := 0; i < h.NrNeurons; i++ {
		for j := 0; j < h.NrNeurons; j++ {
			h.Weights[i][j] += pattern[i] * pattern[j]
			if i == j {
				// no self connection
				h.Weights[i][j] = 0
			}
		}
	}

	return true
}

// NewNet creates a new Hopfield net and returns it.
func NewNet(NrNeurons int) HopfieldNet {
	weights := make([][]float64, NrNeurons)
	for i := range weights {
		weights[i] = make([]float64, NrNeurons)
	}
	return HopfieldNet{weights, NrNeurons}
}

func main() {
	train1 := []float64{1, 1, 1, 1, 1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1}
	train2 := []float64{1, 1, 1, 1, 1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1}
	train3 := []float64{1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1}
	test := []float64{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, 1}
	hNet := NewNet(25)
	hNet.InsertPattern(train1)
	hNet.InsertPattern(train2)
	hNet.InsertPattern(train3)
	//fmt.Print(hNet)

	fmt.Println(hNet)
	fmt.Println(train1)
	fmt.Println(hNet.Recall(test))

}
