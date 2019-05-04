package main

func normalize(membershipValues []float64) []float64 {
	if max, isNorm := maxInSlice(membershipValues); isNorm == false {
		for i, v := range membershipValues {
			membershipValues[i] = v / max
		}
	}
	return membershipValues
}

func normalize2(membership map[float64]float64) map[float64]float64 {
	if max, isNorm := maxInSlice(MapValues(membership)); isNorm == false {
		for k, v := range membership {
			membership[k] = v / max
		}
	}
	return membership
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

func ComparisonTableMembersip(table [][]float64) []float64 {
	n := len(table[0])
	result := make([]float64, n)
	multall2 := make([]float64, 0)
	summAll := 0.0
	//multAll := 1.0

	for i := 0; i < n; i++ {
		tempMult := 1.0
		for j := 0; j < n; j++ {
			tempMult *= table[i][j]
		}
		multall2 = append(multall2, sqrn(tempMult, float64(n)))
		summAll += multall2[i]
	}
	for i := 0; i < n; i++ {
		result[i] = multall2[i] / summAll
	}
	/*
		for k := 0; k < n; k++ {
			for i := 0; i < n; i++ {
				multiplyAll := 1.0
				for j := 0; j < n; j++ {
					multiplyAll *= table[i][j]
				}
				summAll += sqrn(float64(multiplyAll), float64(n))
				multAll *= float64(table[k][i])
			}
			result[k] = sqrn(multAll, float64(n)) / summAll
		}*/
	return result
}
