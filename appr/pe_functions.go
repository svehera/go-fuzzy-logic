package appr

// Triangular represents triangular Pi-approximation membership function
func Triangular(x []float64, a, b, c float64) []float64 {
	y := make([]float64, 0)

	for _, v := range x {
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
	y := make([]float64, 0)

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
