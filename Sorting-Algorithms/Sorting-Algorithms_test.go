package sortingalgorithms_test

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
	"time"

	"sortingAlgortihms/bogo"
	"sortingAlgortihms/bubble"
	"sortingAlgortihms/insertion"
	"sortingAlgortihms/selection"
)

var arrayLengths = []int{10}

func TestSortingAlgorithms(t *testing.T) {
	algorithms := []struct {
		name string
		sort func([]int)
	}{
		{"Bubble", bubble.Sort},
		{"Bogo", bogo.Sort},
		{"Insertion", insertion.Sort},
		{"Selection", selection.Sort},
	}

	for _, n := range arrayLengths {
		t.Run(fmt.Sprintf("Length_%d", n), func(t *testing.T) {
			original := make([]int, n)
			for i := range original {
				original[i] = rand.IntN(20000)
			}

			for _, algorithm := range algorithms {

				t.Run(algorithm.name, func(t *testing.T) {
					arr := slices.Clone(original)

					start := time.Now()
					algorithm.sort(arr)
					elapsed := time.Since(start)

					if !slices.IsSorted(arr) {
						t.Errorf("%s failed to sort an array of the length %d\ngot:%v\noriginal:%v", algorithm.name, n, arr, original)
					}

					t.Logf("Duration: %v", elapsed)
				})
			}
		})
	}
}
