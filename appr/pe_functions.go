package appr

import (
	"fmt"
	"github.com/svehera/go-fuzzy-logic/util"
)

type Triangular string

// Triangular represents triangular Pi-approximation membership function
func (t Triangular) Approximate(support map[float64]float64) []float64 {
	//func Triangular(x []float64, a, b, c float64) []float64 {
	y := make([]float64, 0)
	a, b, c := getParamsTrinagular(support)
	fmt.Printf("a=%f b=%f c=%f\n ", a, b, c)
	//a, b, c := 0.0, 0.0, 0.0

	for _, v := range util.MapKeys(support) {
		if v <= a {
			y = append(y, 0)
		} else if v >= a && v <= b {
			y = append(y, (v-a)/(b-a))
		} else if v >= b && v <= c {
			y = append(y, (c-v)/(c-b))
		} else if c <= v {
			y = append(y, 0)
		}
	}

	return y
}

// Trapezoidal represents trapezoidal Pi-approximation membership function
func Trapezoidal(xs []float64, a, b, c, d float64) []float64 {
	//func Trapezoidal(xs []float64, a, b, c, d float64) []float64 {
	y := make([]float64, 0)
	//a, b, c, d = getParamsTrap(support)
	for _, x := range xs {
		if x <= a {
			y = append(y, 0)
		} else if x >= a && x <= b {
			y = append(y, (x-a)/(b-a))
		} else if x >= b && x <= c {
			y = append(y, 1)
		} else if x >= c && x <= d {
			y = append(y, (d-x)/(d-c))
		} else if d <= x {
			y = append(y, 0)
		}
	}

	return y
}
