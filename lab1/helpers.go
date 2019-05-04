package main

import "sort"

func MapKeys(m map[float64]float64) []float64 {
	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	return keys
}

func MapValues(m map[float64]float64) []float64 {
	var values []float64
	keys := MapKeys(m)
	for _, key := range keys {
		values = append(values, m[key])
	}
	return values
}

func maxInSlice(slice []float64) (float64, bool) {
	sort.Float64s(slice)
	max := slice[len(slice)-1]
	if max == 1.0 {
		return max, true
	}
	return max, false
}
