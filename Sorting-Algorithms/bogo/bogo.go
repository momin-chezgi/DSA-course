package bogo

import (
	"math/rand/v2"
	"time"
)

func Sort(a []int) {
	deadline := time.Now().Add(time.Minute * 3)
	for time.Now().Before(deadline) {
		if is_inceremental(a) {
			break
		}
		rand.Shuffle(len(a), func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})

	}
}

func is_inceremental(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

// var randA []int = a
