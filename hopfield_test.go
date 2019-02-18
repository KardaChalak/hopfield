package hopfield

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewNet(t *testing.T) {
	weightsToString := "[0 0 0 0 0]\n[0 0 0 0 0]\n[0 0 0 0 0]\n[0 0 0 0 0]\n[0 0 0 0 0]\n"
	hNet, err := NewNet(5)
	if err != nil {
		t.Error("Expected no error, got: ", err)
	}
	if hNet.NrNeurons != 5 {
		t.Error("Expected hNet.NrNeurons to be 5, got: ", len(hNet.Weights))
	}
	if hNet.String() != weightsToString {
		t.Error(errorMessage(weightsToString, hNet.String()))

	}

}

func TestInsertPattern(t *testing.T) {
	train1 := []int{1, 1, 1, 1}
	train2 := []int{-1, -1, -1, -1}
	weight1ToString := "[0 1 1 1]\n[1 0 1 1]\n[1 1 0 1]\n[1 1 1 0]\n"
	weight2ToString := "[0 2 2 2]\n[2 0 2 2]\n[2 2 0 2]\n[2 2 2 0]\n"
	hNet, err := NewNet(len(train1))
	if err != nil {
		t.Error("Expected no error but got: ", err)
	}
	err = hNet.InsertPattern(train1)
	if err != nil {
		t.Error("Expected no error for pattern1 but got: ", err)
	}
	if hNet.String() != weight1ToString {
		t.Error(errorMessage(weight1ToString, hNet.String()))
	}

	err = hNet.InsertPattern(train2)
	if err != nil {
		t.Error("Expected no error for pattern2 but got: ", err)
	}

	if hNet.String() != weight2ToString {
		t.Error(errorMessage(weight2ToString, hNet.String()))
	}

}

func TestRecall(t *testing.T) {
	train1 := []int{1, 1, 1, 1}
	train2 := []int{-1, -1, -1, -1}
	distortedPattern1 := []int{-1, 1, 1, 1}
	distortedPattern2 := []int{-1, -1, -1, 1}
	expectedRecall1 := sliceToString([]int{1, 1, 1, 1})
	expectedRecall2 := sliceToString([]int{-1, -1, -1, -1})

	hNet, _ := NewNet(len(train1))
	hNet.InsertPattern(train1)
	hNet.InsertPattern(train2)

	recall1, _ := hNet.Recall(distortedPattern1)
	recall1String := sliceToString(recall1)
	if expectedRecall1 != recall1String {
		t.Error(errorMessage(expectedRecall1, recall1String))
	}

	recall2, _ := hNet.Recall(distortedPattern2)
	recall2String := sliceToString(recall2)
	if expectedRecall2 != recall2String {
		t.Error(errorMessage(expectedRecall2, recall2String))
	}

}

func errorMessage(expected, recivied string) string {
	str := "Expected \n"
	str += expected + " but got :\n" + recivied
	return str
}

func sliceToString(slice []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(slice)), ",")
}
