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
	x := []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1}
	hNet := NewNet(x)
	fmt.Print(hNet)

}
