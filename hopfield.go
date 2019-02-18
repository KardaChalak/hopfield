package main

import (
	"fmt"
	"strings"
)

// HopfieldNet represents the weights between the neurons
// Should be a symmetrical matrix
type HopfieldNet struct {
	Weights [][]float64
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

// NewNet creates a new Hopfield net and returns it.
func NewNet(x []float64) HopfieldNet {
	weights := make([][]float64, len(x))
	for i := range x {
		weights[i] = make([]float64, len(x))
	}
	for i, v := range x {
		for j, w := range x {
			weights[i][j] = v * w

			if i == j {
				weights[i][j] = 0
			}
		}
	}

	return HopfieldNet{weights}
}

func main() {
	train := []float64{1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1}
	test := []float64{1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, 1, -1, -1, 1, -1, -1, -1, -1, -1, 1, 1, 1, -1}
	hNet := NewNet(train)
	//fmt.Print(hNet)

	fmt.Println(hNet)
	fmt.Println(train)
	fmt.Print(hNet.Recall(test))

}
