package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  10,
		"second": 40,
	}

	floats := map[string]float64{
		"first":  0.1,
		"second": 0.4,
	}

	fmt.Printf("Non generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)
}

// SumInts adds together the values of m
func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumFloats adds together the values of m
func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumIntsOrFloats adds together the values of m, supports both int64 or float64
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
