package randomutils_test

import (
	"fmt"
	. "qipai/pkg/randomutils"
	"testing"
)

type testWeight struct {
	Weights  []int
	HasError bool
}

var (
	testWeights = []*testWeight{
		&testWeight{
			Weights:  []int{1, 2, 3},
			HasError: false,
		},
		&testWeight{
			Weights:  []int{0, 0, 0},
			HasError: true,
		},
		&testWeight{
			Weights:  []int{1, 2, -1},
			HasError: true,
		},
	}
)

func TestRandomWeights(t *testing.T) {
	for _, testWeight := range testWeights {
		err := testRandomWeights(testWeight.Weights)
		if err != nil && testWeight.HasError {
			continue
		}
		if err == nil && !testWeight.HasError {
			continue
		}
		t.Fatalf("weight %#v expected has error [%v],but get has error [%v]", testWeight.Weights, testWeight.HasError, err != nil)
	}
}

func testRandomWeights(weights []int) (err error) {
	defer func() {
		terr := recover()
		if terr != nil {
			err = fmt.Errorf("error")
		}
	}()
	RandomWeights(weights)
	return
}
