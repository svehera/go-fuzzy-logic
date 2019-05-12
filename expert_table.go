package fuzzy

type ExpertTable struct {
	FuzzySet    []float64
	Experts     int
	ResultTable [][]float64

	Membership map[float64]float64
	Core       map[float64]float64
	Support    map[float64]float64
	Limit      map[float64]float64
}

func NewExpertTable(experts int, fuzzySet []float64) *ExpertTable {
	return &ExpertTable{Experts: experts, FuzzySet: fuzzySet}
}

func (et *ExpertTable) FillResultTable(estimate func(fuzzySet []float64) []float64) {
	et.ResultTable = make([][]float64, et.Experts)
	for i := range et.ResultTable {
		et.ResultTable[i] = make([]float64, len(et.FuzzySet))
		et.ResultTable[i] = estimate(et.FuzzySet)
	}
}

func (et *ExpertTable) CalculateMembership() {
	et.Membership = make(map[float64]float64)
	for i := 0; i < len(et.FuzzySet); i++ {
		et.Membership[et.FuzzySet[i]] = membershipFunc(et.ResultTable)[i]
	}
}

func membershipFunc(table [][]float64) []float64 {
	result := make([]float64, len(table[0]))

	for i := range result {
		var sum = 0.0
		for j := range table {
			sum += float64(table[j][i])
		}
		result[i] = sum * (1.0 / float64(len(table)))
	}
	return result
}
