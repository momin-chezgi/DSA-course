package sortingalgorithms_test

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strconv"
	"testing"
	"time"

	"sortingAlgortihms/bogo"
	"sortingAlgortihms/bubble"
	"sortingAlgortihms/insertion"
	"sortingAlgortihms/selection"
)

func TestBenchmarkSorting(t *testing.T) {
	sizes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 50, 100, 200, 500, 1000, 10000}

	file, err := os.Create("../results.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"Algorithm", "ArraySize", "AverageTime_ns",
	})

	algorithms := []struct {
		name string
		sort func([]int)
	}{
		{"Bubble", bubble.Sort},
		{"Selection", selection.Sort},
		{"Insertion", insertion.Sort},
		{"Bogo", bogo.Sort},
	}

	for _, size := range sizes {
		original := make([]int, size)
		for i := range original {
			original[i] = rand.Intn(100000)
		}

		for _, algorithm := range algorithms {
			if algorithm.name == "Bogo" && size > 10 {
				continue
			}

			arr := slices.Clone(original)

			start := time.Now()
			algorithm.sort(arr)
			elapsed := time.Since(start)

			if !sort.IntsAreSorted(arr) {
				t.Fatalf("%s failed for size %d",
					algorithm.name, size)
			}

			writer.Write([]string{
				algorithm.name,
				strconv.Itoa(size),
				strconv.FormatInt(elapsed.Nanoseconds(), 10),
			})

			fmt.Printf("%-10s n=%-5d time=%d ns\n",
				algorithm.name, size, elapsed.Nanoseconds())
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Results saved to results.csv")
}
