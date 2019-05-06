package main

type ExpertTable struct {
	Heights     []int
	ExpertsNo   int
	ResultTable [][]int
	Membership  map[float64]float64
	Core        map[float64]float64
	Support     map[float64]float64
	Limit       map[float64]float64
}

func (et *ExpertTable) FillHeights(minHeight, maxHeight, inc int) {
	for h := minHeight; h <= maxHeight; h += inc {
		et.Heights = append(et.Heights, h)
	}
}

func (et *ExpertTable) FillResultTable(estimate func([]int) []int) {
	if len(et.Heights) == 0 {
		et.FillHeights(150, 200, 5)
	}
	et.ResultTable = make([][]int, et.ExpertsNo)
	for i := range et.ResultTable {
		et.ResultTable[i] = estimate(et.Heights)
	}
}

func (et *ExpertTable) CalculateMembershipFunc() {
	membershipFuncValues := CalculateMembership(et.ResultTable)
	et.Membership = make(map[float64]float64)
	for i := 0; i < len(et.Heights); i++ {
		et.Membership[et.HeightsAsFloat()[i]] = membershipFuncValues[i]
	}
}

func (et *ExpertTable) FuzzyCore() {
	et.Core = make(map[float64]float64)

	for k, v := range et.Membership {
		if v == 1 {
			et.Core[k] = v - 0.03
		}
	}
}

func (et *ExpertTable) FuzzySupport() {
	et.Support = make(map[float64]float64)

	for k, v := range et.Membership {
		if v > 0 {
			et.Support[k] = 0.03
		}
	}
}

func (et *ExpertTable) FuzzyLimit() {
	et.Limit = make(map[float64]float64)

	for k, v := range et.Membership {
		if v > 0 && v < 1 {
			et.Limit[k] = v - 0.03
		}
	}
}

func (et *ExpertTable) HeightsAsFloat() []float64 {
	result := make([]float64, len(et.Heights))
	for i := range result {
		result[i] = float64(et.Heights[i])
	}
	return result
}
