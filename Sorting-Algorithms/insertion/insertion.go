package insertion

func Sort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i - 1
		for j >= 0 && a[j] > a[j+1] {
			a[j], a[j+1] = a[j+1], a[j]
			j--
		}
	}
}

// CPU Frequency: 2GHz
