package selection

func Sort(a []int) {
	for i := 0; i < len(a); i++ {
		minIndex := i
		for j := i; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

// CPU Frequency: 2GHz
