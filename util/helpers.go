package util

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrMinGreaterMax = errors.New("min greater than max")
	ErrMinZeroValue  = errors.New("min equal zero")
)

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

func MaxInSlice(slice []float64) float64 {
	sort.Float64s(slice)
	return slice[len(slice)-1]
}

func MinInSlice(slice []float64) float64 {
	sort.Float64s(slice)
	return slice[0]
}

func SetHeights(minHeight, maxHeight, increment uint8) (heights []float64, err error) {
	if minHeight == 0 {
		return nil, ErrMinZeroValue
	}
	if minHeight > maxHeight {
		return nil, ErrMinGreaterMax
	}
	for h := minHeight; h <= maxHeight; h += increment {
		heights = append(heights, float64(h))
	}
	return
}

func SliceAsTable(table [][]float64) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[0]); j++ {
			fmt.Printf("%f  ", table[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
