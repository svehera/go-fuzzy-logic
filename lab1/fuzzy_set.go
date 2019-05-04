package main

func normalize(membershipValues []float64) []float64 {
	if max, isNorm := maxInSlice(membershipValues); isNorm == false {
		for i, v := range membershipValues {
			membershipValues[i] = v / max
		}
	}
	return membershipValues
}

func CalculateMembership(table [][]int) []float64 {

	result := make([]float64, len(table[0]))
	for i := range result {
		aff := func() float64 {
			var sum = 0.0
			for j := range table {
				sum += float64(table[j][i])
			}
			return sum
		}
		result[i] = aff() * (1.0 / float64(len(table)))
	}
	return result
}
